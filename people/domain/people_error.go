/*
 * File: people_error.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrPersonNotFoundCode           = "ERR_PERSON_NOT_FOUND"
	ErrPersonUserIdAlreadyExistCode = "ERR_PERSON_USER_ID_ALREADY_EXIST"
	ErrPersonIdAlreadyDeletedCode   = "ERR_PERSON_ID_ALREADY_DELETED"
)

var (
	ErrPersonNotFound = errDomain.NewErr().
				SetCode(ErrPersonNotFoundCode).
				SetDescription("PERSON NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("UpdatePerson")

	ErrPersonUserIdAlreadyExist = errDomain.NewErr().
					SetCode(ErrPersonUserIdAlreadyExistCode).
					SetDescription("USER ID ALREADY EXIST").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreatePerson")

	ErrPersonIdAlreadyDeleted = errDomain.NewErr().
					SetCode(ErrPersonIdAlreadyDeletedCode).
					SetDescription("PERSON ID ALREADY DELETED").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("DeletePerson")
)
