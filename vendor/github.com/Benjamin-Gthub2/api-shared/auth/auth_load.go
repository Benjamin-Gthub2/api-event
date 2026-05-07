/*
 * File: auth_load.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Load Auth Middleware
 *
 * Last Modified: 2023-11-27
 */

package auth

import (
	"time"

	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	auth "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	authRest "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	authUseCase "github.com/Benjamin-Gthub2/api-shared/auth/usecase"
)

func LoadAuthMiddleware() auth.AuthMiddleware {
	timeoutContext := time.Duration(60) * time.Second
	authJWTRepository := authRepository.NewAuthRepository()
	authUCase := authUseCase.NewAuthUseCase(authJWTRepository, timeoutContext)
	authMiddleware := authRest.NewAuthMiddleware(authUCase)
	return authMiddleware
}
