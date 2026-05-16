/*
 * File: setup_registrations.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the setup of the registrations certificate module.
 *
 * Last Modified: 2026-05-16
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"

	eventsSharedRepository "github.com/Benjamin-Gthub2/api-event/events-shared/infrastructure/persistence/mysql"
	registrationCertificatePdfRepository "github.com/Benjamin-Gthub2/api-event/registrations-certificate/infrastructure/reports/pdf/registrations"
	registrationCertificateHttpDelivery "github.com/Benjamin-Gthub2/api-event/registrations-certificate/interfaces/rest"
	registrationCertificateUseCase "github.com/Benjamin-Gthub2/api-event/registrations-certificate/usecase"
	registrationsRepository "github.com/Benjamin-Gthub2/api-event/registrations/infrastructure/persistence/mysql"
)

func LoadRegistrationsCertificate(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()

	eventSharedRepository := eventsSharedRepository.NewEventSharedRepository(clock, 60)
	registrationRepo := registrationsRepository.NewRegistrationsRepository(clock, 60, eventSharedRepository)
	certificatePdfRepo := registrationCertificatePdfRepository.NewRegistrationCertificatePdfRepository(clock, 60)

	certificateUCase := registrationCertificateUseCase.NewRegistrationsCertificateUseCase(
		registrationRepo,
		certificatePdfRepo,
		timeoutContext,
	)

	_ = authJWTRepository
	registrationCertificateHttpDelivery.NewRegistrationsCertificateHandler(certificateUCase, router, authMiddleware)
}
