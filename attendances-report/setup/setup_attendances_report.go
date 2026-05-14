package setup

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Benjamin-Gthub2/api-shared/auth"
	authRepository "github.com/Benjamin-Gthub2/api-shared/auth/infrastructure/jwt"
	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"

	eventsSharedRepository "github.com/Benjamin-Gthub2/api-event/events-shared/infrastructure/persistence/mysql"
	attendancesRepository "github.com/Benjamin-Gthub2/api-event/attendances/infrastructure/persistence/mysql"

	attendancesReportXlsxRepository "github.com/Benjamin-Gthub2/api-event/attendances-report/infrastructure/reports/xlsx"
	attendancesReportHttpDelivery "github.com/Benjamin-Gthub2/api-event/attendances-report/interfaces/rest"
	attendancesReportUseCase "github.com/Benjamin-Gthub2/api-event/attendances-report/usecase"
)

func LoadAttendancesReport(router *gin.Engine) {
	timeoutContext := time.Duration(60) * time.Second
	clock := smartClock.NewClock()
	authJWTRepository := authRepository.NewAuthRepository()
	eventSharedRepo := eventsSharedRepository.NewEventSharedRepository(clock, 60)
	attendanceRepo := attendancesRepository.NewAttendancesRepository(clock, 60, eventSharedRepo)
	xlsxRepo := attendancesReportXlsxRepository.NewAttendancesReportXlsxRepository(clock, 60)
	authMw := auth.LoadAuthMiddleware()

	attendancesReportUCase := attendancesReportUseCase.NewAttendancesReportUseCase(
		xlsxRepo,
		attendanceRepo,
		authJWTRepository,
		timeoutContext,
	)
	attendancesReportHttpDelivery.NewAttendancesReportHandler(attendancesReportUCase, router, authMw)
}
