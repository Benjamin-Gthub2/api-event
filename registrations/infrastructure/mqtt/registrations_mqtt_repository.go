/*
 * File: registrations_mqtt_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the mqtt repository for registrations.
 *
 * Last Modified: 2026-04-22
 */

package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsMqttRepository struct {
	client *mqtt.Client
	err    *errDomain.SmartError
}

func NewRegistrationsRTRepository(
	client *mqtt.Client,
) registrationsDomain.RegistrationsRTRepository {
	rep := &registrationsMqttRepository{
		client: client,
		err:    errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	return rep
}
