package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	registrationStatusesRepository "github.com/Benjamin-Gthub2/api-event/registration-statuses/infrastructure/persistence/mysql"
	registrationStatusesHttpDelivery "github.com/Benjamin-Gthub2/api-event/registration-statuses/interfaces/rest"
	registrationStatusesUseCase "github.com/Benjamin-Gthub2/api-event/registration-statuses/usecase"
)

func LoadRegistrationStatuses(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	registrationStatusRepository := registrationStatusesRepository.NewRegistrationStatusesRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	registrationStatusesUCase := registrationStatusesUseCase.NewRegistrationStatusesUseCase(
		registrationStatusRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	registrationStatusesHttpDelivery.NewRegistrationStatusesHandler(registrationStatusesUCase, router, authMiddleware)
}
