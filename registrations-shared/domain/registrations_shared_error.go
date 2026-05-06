package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrRegistrationStatusNotFoundCode = "ERR_REGISTRATION_STATUS_NOT_FOUND"
)

var (
	ErrRegistrationStatusNotFound = errDomain.NewErr().
		SetCode(ErrRegistrationStatusNotFoundCode).
		SetDescription("REGISTRATION STATUS NOT FOUND").
		SetLevel(errDomain.LevelError).
		SetHttpStatus(http.StatusNotFound).
		SetLayer(errDomain.UseCase).
		SetFunction("GetStatusByCode")
)
