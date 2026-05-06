/*
 * File: auth_usecase.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Initializing use cases for auth
 *
 * Last Modified: 2023-11-26
 */

package usecase

import (
	"time"

	"github.com/Benjamin-Gthub2/api-shared/auth/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(
	authRepository domain.AuthRepository,
	timeout time.Duration,
) domain.AuthUseCase {
	return &authUseCase{
		authRepository: authRepository,
		contextTimeout: timeout,
	}
}
