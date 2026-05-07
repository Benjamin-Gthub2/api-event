/*
 * File: setup_user_types.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the user types.
 *
 * Last Modified: 2023-12-28
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

	userTypesRepository "github.com/Benjamin-Gthub2/api-event/user-types/infrastructure/persistence/mysql"
	userTypesHttpDelivery "github.com/Benjamin-Gthub2/api-event/user-types/interfaces/rest"
	userTypesUseCase "github.com/Benjamin-Gthub2/api-event/user-types/usecase"
)

func LoadUserTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	userTypeRepository := userTypesRepository.NewUserTypesRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	userTypesUCase := userTypesUseCase.NewUserTypesUseCase(
		userTypeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	userTypesHttpDelivery.NewUserTypesHandler(userTypesUCase, router, authMiddleware)
}
