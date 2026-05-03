package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	materialsIssuedRepository "github.com/Benjamin-Gthub2/api-event/materials-issued/infrastructure/persistence/mysql"
	materialsIssuedHttpDelivery "github.com/Benjamin-Gthub2/api-event/materials-issued/interfaces/rest"
	materialsIssuedUseCase "github.com/Benjamin-Gthub2/api-event/materials-issued/usecase"
)

func LoadMaterialsIssued(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	materialIssuedRepository := materialsIssuedRepository.NewMaterialsIssuedRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	materialsIssuedUCase := materialsIssuedUseCase.NewMaterialsIssuedUseCase(
		materialIssuedRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	materialsIssuedHttpDelivery.NewMaterialsIssuedHandler(materialsIssuedUCase, router, authMiddleware)
}
