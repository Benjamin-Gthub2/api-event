/*
 * File: setup_people.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the microservice people.
 *
 * Last Modified: 2023-12-28
 */

package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	peopleRepository "github.com/Benjamin-Gthub2/api-event/people/infrastructure/persistence/mysql"
	peopleHttpDelivery "github.com/Benjamin-Gthub2/api-event/people/interfaces/rest"
	peopleUseCase "github.com/Benjamin-Gthub2/api-event/people/usecase"
)

func LoadPeople(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	Repository := peopleRepository.NewPeopleRepository(clock, 60)
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	peopleUCase := peopleUseCase.NewPeopleUseCase(
		Repository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	peopleHttpDelivery.NewPeopleHandler(peopleUCase, router, authMiddleware)
}
