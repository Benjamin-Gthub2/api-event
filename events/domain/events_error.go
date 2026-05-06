/*
 * File: events_error.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the modules errors.
 *
 * Last Modified: 2026-04-15
 */

package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrEventNotFoundCode             = "ERR_EVENT_NOT_FOUND"
	ErrEventDocumentAlreadyExistCode = "ERR_EVENT_DOCUMENT_ALREADY_EXIST"
	ErrEventIdHasBeenDeletedCode     = "ERR_EVENT_ID_HAS_BEEN_DELETED"
)

var (
	ErrEventNotFound = errDomain.NewErr().
				SetCode(ErrEventNotFoundCode).
				SetDescription("EVENT NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdateEvent")

	ErrEventDocumentAlreadyExist = errDomain.NewErr().
					SetCode(ErrEventDocumentAlreadyExistCode).
					SetDescription("DOCUMENT ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateEvent")

	ErrEventIdHasBeenDeleted = errDomain.NewErr().
					SetCode(ErrEventIdHasBeenDeletedCode).
					SetDescription("ID HAS BEEN DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeleteAccount")

	ErrEventCodeAlreadyExist = errDomain.NewErr().
					SetCode(ErrEventDocumentAlreadyExistCode).
					SetDescription("CODE ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateEvent")
)
