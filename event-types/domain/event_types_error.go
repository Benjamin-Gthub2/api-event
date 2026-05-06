package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrEventTypeNotFoundCode = "ERR_EVENT_TYPE_NOT_FOUND"
)

var (
	ErrEventTypeNotFound = errDomain.NewErr().
				SetCode(ErrEventTypeNotFoundCode).
				SetDescription("EVENT TYPE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetEventTypeById")
)
