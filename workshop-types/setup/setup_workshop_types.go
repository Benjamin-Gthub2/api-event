package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	workshopTypesRepository "github.com/Benjamin-Gthub2/api-event/workshop-types/infrastructure/persistence/mysql"
	workshopTypesHttpDelivery "github.com/Benjamin-Gthub2/api-event/workshop-types/interfaces/rest"
	workshopTypesUseCase "github.com/Benjamin-Gthub2/api-event/workshop-types/usecase"
)

func LoadWorkshopTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	workshopTypeRepository := workshopTypesRepository.NewWorkshopTypesRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	workshopTypesUCase := workshopTypesUseCase.NewWorkshopTypesUseCase(
		workshopTypeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	workshopTypesHttpDelivery.NewWorkshopTypesHandler(workshopTypesUCase, router, authMiddleware)
}
