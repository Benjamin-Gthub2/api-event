package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrAttendanceNotFoundCode = "ERR_ATTENDANCE_NOT_FOUND"
)

var (
	ErrAttendanceNotFound = errDomain.NewErr().
				SetCode(ErrAttendanceNotFoundCode).
				SetDescription("ATTENDANCE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetAttendanceById")
)
