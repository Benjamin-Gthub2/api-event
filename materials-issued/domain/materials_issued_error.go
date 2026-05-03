package domain

import (
	"net/http"

	errDomain "github.com/smart0n3/api-shared/error-core/domain"
)

const (
	ErrMaterialIssuedNotFoundCode = "ERR_MATERIAL_ISSUED_NOT_FOUND"
)

var (
	ErrMaterialIssuedNotFound = errDomain.NewErr().
					SetCode(ErrMaterialIssuedNotFoundCode).
					SetDescription("MATERIAL ISSUED NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetMaterialIssuedById")
)
