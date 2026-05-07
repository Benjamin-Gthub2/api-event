package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrStoreMiddleNotFoundCode         = "ERR_STORES_MIDDLE_NOT_FOUND"
	ErrUserInsufficientPermissionsCode = "ERR_USER_INSUFFICIENT_PERMISSIONS"
)

var (
	ErrGetStoresMiddleNotFound = errDomain.NewErr().
					SetCode(ErrStoreMiddleNotFoundCode).
					SetDescription("STORES MIDDLE NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.Infra).
					SetFunction("GetStrStoreIds")

	ErrUserInsufficientPermissions = errDomain.NewErr().
					SetCode(ErrUserInsufficientPermissionsCode).
					SetDescription("USER INSUFFICIENT PERMISSIONS").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.Infra).
					SetFunction("GetStrStoreIds")
)
