/*
 * File: auth_jwt_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Repository for auth.
 *
 * Last Modified: 2023-11-26
 */

package jwt

import (
	"github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

type authJWTRepo struct {
	err *errDomain.SmartError
}

func NewAuthRepository() domain.AuthRepository {
	rep := &authJWTRepo{
		err: errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
