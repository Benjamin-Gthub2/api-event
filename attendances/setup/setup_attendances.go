package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	"github.com/Benjamin-Gthub2/api-shared/mqtt"
	validationsRepository "github.com/Benjamin-Gthub2/api-shared/validations/infrastructure/persistence/mysql"

	eventsSharedRepository "github.com/Benjamin-Gthub2/api-event/events-shared/infrastructure/persistence/mysql"

	attendancesMqttRepository "github.com/Benjamin-Gthub2/api-event/attendances/infrastructure/mqtt"
	attendancesRepository "github.com/Benjamin-Gthub2/api-event/attendances/infrastructure/persistence/mysql"
	attendancesHttpDelivery "github.com/Benjamin-Gthub2/api-event/attendances/interfaces/rest"
	attendancesUseCase "github.com/Benjamin-Gthub2/api-event/attendances/usecase"
)

func LoadAttendances(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	validationRepository := validationsRepository.NewValidationsRepository(60)
	authJWTRepository := authRepository.NewAuthRepository()
	eventSharedRepository := eventsSharedRepository.NewEventSharedRepository(clock, 60)
	attendanceRepository := attendancesRepository.NewAttendancesRepository(clock, 60, eventSharedRepository)
	authMiddleware := auth.LoadAuthMiddleware()
	attendanceMqttRepository := attendancesMqttRepository.NewAttendancesRTRepository(mqtt.MqttClient)

	attendancesUCase := attendancesUseCase.NewAttendancesUseCase(
		attendanceRepository,
		attendanceMqttRepository,
		validationRepository,
		authJWTRepository,
		timeoutContext,
	)
	attendancesHttpDelivery.NewAttendancesHandler(attendancesUCase, router, authMiddleware)
}
