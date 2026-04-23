package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	sessionsRepository "github.com/Benjamin-Gthub2/api-event/sessions/infrastructure/persistence/mysql"
	sessionsHttpDelivery "github.com/Benjamin-Gthub2/api-event/sessions/interfaces/rest"
	sessionsUseCase "github.com/Benjamin-Gthub2/api-event/sessions/usecase"
)

func LoadSessions(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	sessionRepository := sessionsRepository.NewSessionsRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	sessionsUCase := sessionsUseCase.NewSessionsUseCase(
		sessionRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	sessionsHttpDelivery.NewSessionsHandler(sessionsUCase, router, authMiddleware)
}
