/*
 * File: main.go
 * Author: Benjamin
 * Copyright: 2026, Benjamin Alexander.
 * License: MIT
 *
 * Purpose:
 * This is file content the main of the microservice event.
 *
 * Last Modified: 2023-12-28
 */

package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/config"
	"github.com/Benjamin-Gthub2/api-shared/db"
	"github.com/Benjamin-Gthub2/api-shared/mqtt"

	attendancesSetup "github.com/Benjamin-Gthub2/api-event/attendances/setup"
	eventTypesSetup "github.com/Benjamin-Gthub2/api-event/event-types/setup"
	eventsSetup "github.com/Benjamin-Gthub2/api-event/events/setup"
	materialsIssuedSetup "github.com/Benjamin-Gthub2/api-event/materials-issued/setup"
	peopleSetup "github.com/Benjamin-Gthub2/api-event/people/setup"
	registrationStatusesSetup "github.com/Benjamin-Gthub2/api-event/registration-statuses/setup"
	registrationsCertificateSetup "github.com/Benjamin-Gthub2/api-event/registrations-certificate/setup"
	registrationsSetup "github.com/Benjamin-Gthub2/api-event/registrations/setup"
	sessionsSetup "github.com/Benjamin-Gthub2/api-event/sessions/setup"
	usersSetup "github.com/Benjamin-Gthub2/api-event/users/setup"
	workshopSpeakersSetup "github.com/Benjamin-Gthub2/api-event/workshop-speakers/setup"
	workshopTypesSetup "github.com/Benjamin-Gthub2/api-event/workshop-types/setup"
	workshopsSetup "github.com/Benjamin-Gthub2/api-event/workshops/setup"
)

func main() {
	cfg := config.Configuration{
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DB: config.DB{
			DbDatabase: os.Getenv("DB_DATABASE"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		LoggingUrl: os.Getenv("LOGGING_URL"),
	}
	err := db.InitClients(cfg)
	if err != nil {
		return
	}
	defer db.Disconnect()

	_, err = mqtt.ConnectToMQTT()
	if err != nil {
		fmt.Println("MQTT connection failed, continuing without it:", err)
		//return
	}

	router := gin.Default()

	corsOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*"
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Tenant-Id"},
		ExposeHeaders:    []string{"X-Tenant-Id"},
		AllowCredentials: corsOrigin != "*",
	}))

	router.GET("/health", func(c *gin.Context) { c.Status(200) })

	usersSetup.LoadUsers(router)
	registrationsSetup.LoadRegistrations(router)
	peopleSetup.LoadPeople(router)
	workshopsSetup.LoadWorkshops(router)
	workshopSpeakersSetup.LoadWorkshopSpeakers(router)
	sessionsSetup.LoadSessions(router)
	workshopTypesSetup.LoadWorkshopTypes(router)
	eventTypesSetup.LoadEventTypes(router)
	eventsSetup.LoadEvents(router)
	attendancesSetup.LoadAttendances(router)
	materialsIssuedSetup.LoadMaterialsIssued(router)
	registrationStatusesSetup.LoadRegistrationStatuses(router)
	registrationsCertificateSetup.LoadRegistrationsCertificate(router)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	serverPort := fmt.Sprintf(":%s", port)
	err = router.Run(serverPort)
	if err != nil {
		return
	}

}
