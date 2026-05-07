package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrSessionNotFoundCode = "ERR_SESSION_NOT_FOUND"
)

var (
	ErrSessionNotFound = errDomain.NewErr().
				SetCode(ErrSessionNotFoundCode).
				SetDescription("SESSION NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetSessionById")
)
