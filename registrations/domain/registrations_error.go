/*
 * File: registrations_error.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the errors of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrRegistrationsNotFoundCode = "ERR_REGISTRATION_NOT_FOUND"
)

var (
	ErrRegistrationsNotFound = errDomain.NewErr().
					SetCode(ErrRegistrationsNotFoundCode).
					SetDescription("REGISTRATION NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetQrRegistrationById")

	ErrUseCaseRegistrationsNotFound = errDomain.NewErr().
					SetCode(ErrRegistrationsNotFoundCode).
					SetDescription("TRANSFER RECEIPT NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetRegistrationById")
)
