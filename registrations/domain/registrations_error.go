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

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrRegistrationsNotFoundCode      = "ERR_REGISTRATION_NOT_FOUND"
	ErrEventsNotFoundCode             = "ERR_EVENT_NOT_FOUND"
	ErrPeopleNotFoundCode             = "ERR_PERSON_NOT_FOUND"
	ErrChangeOfStatusIsNotAllowedCode = "ERR_CHANGE_OF_STATUS_IS_NOT_ALLOWED"
)

var (
	ErrRegistrationsNotFound = errDomain.NewErr().
					SetCode(ErrRegistrationsNotFoundCode).
					SetDescription("REGISTRATION NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetQrRegistrationById")
	ErrEventNotFound = errDomain.NewErr().
				SetCode(ErrEventsNotFoundCode).
				SetDescription("EVENT NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateRegistration")
	ErrPersonNotFound = errDomain.NewErr().
				SetCode(ErrEventsNotFoundCode).
				SetDescription("PERSON NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateRegistration")
	ErrUseCaseRegistrationsNotFound = errDomain.NewErr().
					SetCode(ErrRegistrationsNotFoundCode).
					SetDescription("TRANSFER RECEIPT NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetRegistrationById")
	ErrChangeOfStatusIsNotAllowed = errDomain.NewErr().
					SetCode(ErrChangeOfStatusIsNotAllowedCode).
					SetDescription("CHANGE OF STATUS IS NOT POSSIBLE").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("UpdateRegistrationApprovalStatus")
)
