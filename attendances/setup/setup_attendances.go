package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/smart0n3/api-shared/auth"
	authRepository "github.com/smart0n3/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/smart0n3/api-shared/clock"
	validationsRepository "github.com/smart0n3/api-shared/validations/infrastructure/persistence/mysql"

	attendancesRepository "github.com/Benjamin-Gthub2/api-event/attendances/infrastructure/persistence/mysql"
	attendancesHttpDelivery "github.com/Benjamin-Gthub2/api-event/attendances/interfaces/rest"
	attendancesUseCase "github.com/Benjamin-Gthub2/api-event/attendances/usecase"
)

func LoadAttendances(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	attendanceRepository := attendancesRepository.NewAttendancesRepository(clock, 60)
	authMiddleware := auth.LoadAuthMiddleware()

	attendancesUCase := attendancesUseCase.NewAttendancesUseCase(
		attendanceRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	attendancesHttpDelivery.NewAttendancesHandler(attendancesUCase, router, authMiddleware)
}
