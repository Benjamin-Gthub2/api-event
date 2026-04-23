package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrWorkshopNotFoundCode = "ERR_WORKSHOP_NOT_FOUND"
)

var (
	ErrWorkshopNotFound = errDomain.NewErr().
				SetCode(ErrWorkshopNotFoundCode).
				SetDescription("WORKSHOP NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetWorkshopById")
)
