/*
 * File: setup_users.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is file content the setup of the users.
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

	usersRepository "github.com/Benjamin-Gthub2/api-event/users/infrastructure/persistence/mysql"
	usersHttpDelivery "github.com/Benjamin-Gthub2/api-event/users/interfaces/rest"
	usersUseCase "github.com/Benjamin-Gthub2/api-event/users/usecase"
)

func LoadUsers(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	validationRepository := validationsRepository.NewValidationsRepository(60)
	clock := smartClock.NewClock()
	userRepository := usersRepository.NewUsersRepository(clock, 60)
	authJWTRepository := authRepository.NewAuthRepository()
	authMiddleware := auth.LoadAuthMiddleware()
	usersUCase := usersUseCase.NewUsersUseCase(
		userRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext)
	usersHttpDelivery.NewUsersHandler(usersUCase, router, authMiddleware)
}
