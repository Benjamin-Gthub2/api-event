package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	eventTypesRepository "github.com/Benjamin-Gthub2/api-event/event-types/infrastructure/persistence/mysql"
	eventTypesHttpDelivery "github.com/Benjamin-Gthub2/api-event/event-types/interfaces/rest"
	eventTypesUseCase "github.com/Benjamin-Gthub2/api-event/event-types/usecase"
)

func LoadEventTypes(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	eventTypeRepository := eventTypesRepository.NewEventTypesRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	eventTypesUCase := eventTypesUseCase.NewEventTypesUseCase(
		eventTypeRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	eventTypesHttpDelivery.NewEventTypesHandler(eventTypesUCase, router, authMiddleware)
}
