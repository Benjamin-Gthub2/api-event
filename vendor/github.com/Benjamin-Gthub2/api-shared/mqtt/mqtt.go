package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

//go:embed ca_certificate.pem
var CertPem string

var (
	MqttClient *mqtt.Client
)

func ConnectToMQTT() (client mqtt.Client, err error) {
	mqttPortEnv := os.Getenv("MQTT_PORT")
	mqttHostEnv := os.Getenv("MQTT_HOST")
	mqttPort, err := strconv.Atoi(mqttPortEnv)
	if err != nil {
		fmt.Println("Error al conectar al broker MQTT:", err)
		return nil, err
	}
	mqttURL := fmt.Sprintf("wss://%s:%d/mqtt", mqttHostEnv, mqttPort)
	fmt.Println("mqttUrl", mqttURL)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttURL)
	opts.SetClientID("sigetran_client_" + uuid.New().String())
	opts.SetPingTimeout(10 * time.Second)
	opts.SetKeepAlive(10 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)
	opts.SetCleanSession(false)

	opts.SetOnConnectHandler(onConnectHandler)
	opts.SetConnectionLostHandler(connectionLostHandler)
	opts.SetReconnectingHandler(connectionReconnectionHandler)

	tlsConfig := NewTlsConfig()
	opts.SetTLSConfig(tlsConfig)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error al conectar al broker MQTT:", token.Error())
		return nil, token.Error()
	}

	MqttClient = &client

	return client, nil
}

func onConnectHandler(client mqtt.Client) {
	fmt.Println("Connect broker!!")
	MqttClient = &client
}

func connectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func connectionReconnectionHandler(client mqtt.Client, options *mqtt.ClientOptions) {
	fmt.Println("...... mqtt reconnecting ......")
}

func NewTlsConfig() *tls.Config {
	certpool := x509.NewCertPool()
	ca := []byte(CertPem)
	certpool.AppendCertsFromPEM(ca)
	return &tls.Config{
		RootCAs:            certpool,
		InsecureSkipVerify: true,
	}
}
