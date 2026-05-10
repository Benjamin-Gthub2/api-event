/*
 * File: attendances_mqtt_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the mqtt repository for attendances.
 *
 * Last Modified: 2026-04-22
 */

package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendancesMqttRepository struct {
	client *mqtt.Client
	err    *errDomain.SmartError
}

func NewAttendancesRTRepository(
	client *mqtt.Client,
) attendancesDomain.AttendancesRTRepository {
	rep := &attendancesMqttRepository{
		client: client,
		err:    errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	return rep
}
