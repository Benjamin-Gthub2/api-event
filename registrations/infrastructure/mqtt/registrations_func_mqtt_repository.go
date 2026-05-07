/*
 * File: registrations_func_mqtt_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the mqtt repository functions for registrations.
 *
 * Last Modified: 2026-04-22
 */

package mqtt

import (
	"context"
	"fmt"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

func (r registrationsMqttRepository) SendNotification(
	ctx context.Context,
	topic string,
	payload string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	if r.client == nil {
		return r.err.Clone().SetFunction("SendNotification").SetRaw(fmt.Errorf("MQTT client is not initialized"))
	}

	if !(*r.client).IsConnected() {
		return r.err.Clone().SetFunction("SendNotification").SetRaw(fmt.Errorf("MQTT client is not connected"))
	}

	publish := (*r.client).Publish(topic, 0, false, payload)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-publish.Done():
		if publish.Error() != nil {
			return fmt.Errorf("failed to publish to topic %s: %v", topic, publish.Error())
		}
	}
	return nil
}
