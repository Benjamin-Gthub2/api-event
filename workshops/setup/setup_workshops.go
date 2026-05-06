package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	workshopsRepository "github.com/Benjamin-Gthub2/api-event/workshops/infrastructure/persistence/mysql"
	workshopsHttpDelivery "github.com/Benjamin-Gthub2/api-event/workshops/interfaces/rest"
	workshopsUseCase "github.com/Benjamin-Gthub2/api-event/workshops/usecase"
)

func LoadWorkshops(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	workshopRepository := workshopsRepository.NewWorkshopsRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	workshopsUCase := workshopsUseCase.NewWorkshopsUseCase(
		workshopRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	workshopsHttpDelivery.NewWorkshopsHandler(workshopsUCase, router, authMiddleware)
}
