package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	workshopSpeakersRepository "github.com/Benjamin-Gthub2/api-event/workshop-speakers/infrastructure/persistence/mysql"
	workshopSpeakersHttpDelivery "github.com/Benjamin-Gthub2/api-event/workshop-speakers/interfaces/rest"
	workshopSpeakersUseCase "github.com/Benjamin-Gthub2/api-event/workshop-speakers/usecase"
)

func LoadWorkshopSpeakers(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	workshopSpeakerRepository := workshopSpeakersRepository.NewWorkshopSpeakersRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	workshopSpeakersUCase := workshopSpeakersUseCase.NewWorkshopSpeakersUseCase(
		workshopSpeakerRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	workshopSpeakersHttpDelivery.NewWorkshopSpeakersHandler(workshopSpeakersUCase, router, authMiddleware)
}
