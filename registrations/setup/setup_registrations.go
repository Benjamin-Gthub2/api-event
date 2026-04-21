/*
 * File: setup_registrations.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the setup of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	registrationsRepository "github.com/Benjamin-Gthub2/api-event/registrations/infrastructure/persistence/mysql"
	registrationsHttpDelivery "github.com/Benjamin-Gthub2/api-event/registrations/interfaces/rest"
	registrationsUseCase "github.com/Benjamin-Gthub2/api-event/registrations/usecase"
)

func LoadRegistrations(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	registrationRepository := registrationsRepository.NewRegistrationsRepository(
		clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()
	registrationsUCase := registrationsUseCase.NewRegistrationsUseCase(
		registrationRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	registrationsHttpDelivery.NewRegistrationsHandler(registrationsUCase, router, authMiddleware)
}
