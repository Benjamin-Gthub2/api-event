package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrWorkshopTypeNotFoundCode = "ERR_WORKSHOP_TYPE_NOT_FOUND"
)

var (
	ErrWorkshopTypeNotFound = errDomain.NewErr().
				SetCode(ErrWorkshopTypeNotFoundCode).
				SetDescription("WORKSHOP TYPE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetWorkshopTypeById")
)
