/*
 * File: events_shared_error.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the error in the domain.
 *
 * Last Modified: 2026-04-24
 */

package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrEventNotFoundCode    = "ERR_EVENT_NOT_FOUND"
	ErrWorkshopNotFoundCode = "ERR_WORKSHOP_NOT_FOUND"
	ErrSessionNotFoundCode  = "ERR_SESSION_NOT_FOUND"
)

var (
	ErrEventNotFound = errDomain.NewErr().
				SetCode(ErrEventNotFoundCode).
				SetDescription("EVENT NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.Infra).
				SetFunction("GetEventStatusIdByStatuCode")

	ErrWorkshopNotFound = errDomain.NewErr().
				SetCode(ErrWorkshopNotFoundCode).
				SetDescription("WORKSHOP NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.Infra).
				SetFunction("GetEventAmountTotalById")

	ErrSessionNotFound = errDomain.NewErr().
				SetCode(ErrSessionNotFoundCode).
				SetDescription("SESSION NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.Infra).
				SetFunction("GetEventAmountTotalById")
)
