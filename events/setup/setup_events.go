/*
 * File: setup_events.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the events.
 *
 * Last Modified: 2026-04-15
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	eventsRepository "github.com/Benjamin-Gthub2/api-event/events/infrastructure/persistence/mysql"
	eventsHttpDelivery "github.com/Benjamin-Gthub2/api-event/events/interfaces/rest"
	eventsUseCase "github.com/Benjamin-Gthub2/api-event/events/usecase"
)

func LoadEvents(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	eventRepository := eventsRepository.NewEventsRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	eventUCase := eventsUseCase.NewEventsUseCase(
		eventRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	eventsHttpDelivery.NewEventsHandler(eventUCase, router, authMiddleware)
}
