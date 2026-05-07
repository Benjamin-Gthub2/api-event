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
	"github.com/Benjamin-Gthub2/api-shared/mqtt"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	eventsSharedRepository "github.com/Benjamin-Gthub2/api-event/events-shared/infrastructure/persistence/mysql"
	registrationSharedRepository "github.com/Benjamin-Gthub2/api-event/registrations-shared/infrastructure/persistence/mysql"
	registrationsMqttRepository "github.com/Benjamin-Gthub2/api-event/registrations/infrastructure/mqtt"
	registrationsRepository "github.com/Benjamin-Gthub2/api-event/registrations/infrastructure/persistence/mysql"
	registrationsHttpDelivery "github.com/Benjamin-Gthub2/api-event/registrations/interfaces/rest"
	registrationsUseCase "github.com/Benjamin-Gthub2/api-event/registrations/usecase"
)

func LoadRegistrations(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	eventSharedRepository := eventsSharedRepository.NewEventSharedRepository(clock, 60)
	registrationsSharedRepository := registrationSharedRepository.NewRegistrationSharedRepository(clock, 60)
	registrationRepository := registrationsRepository.NewRegistrationsRepository(
		clock, 60, eventSharedRepository)
	authMiddleware := auth.LoadAuthMiddleware()
	registrationMqttRepository := registrationsMqttRepository.NewRegistrationsRTRepository(mqtt.MqttClient)

	registrationsUCase := registrationsUseCase.NewRegistrationsUseCase(
		registrationRepository,
		registrationMqttRepository,
		registrationsSharedRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	registrationsHttpDelivery.NewRegistrationsHandler(registrationsUCase, router, authMiddleware)
}
