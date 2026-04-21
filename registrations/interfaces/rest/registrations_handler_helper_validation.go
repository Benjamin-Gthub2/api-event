/*
 * File: registrations_handler_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity in the handler.
 *
 * Last Modified: 2026-04-21
 */

package rest

import (
	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationByIdResult struct {
	Data   *registrationsDomain.Registration `json:"data"`
	Status int                               `json:"status" binding:"required"`
}
