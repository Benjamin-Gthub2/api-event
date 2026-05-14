package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
)

type attendancesReportHandler struct {
	attendancesReportUseCase domain.AttendancesReportUseCase
	authMiddleware           authMiddleware.AuthMiddleware
	err                      *errDomain.SmartError
}

func NewAttendancesReportHandler(
	attendancesReportUseCase domain.AttendancesReportUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &attendancesReportHandler{
		attendancesReportUseCase: attendancesReportUseCase,
		authMiddleware:           authMiddleware,
		err:                      errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/attendances/xlsx_report", handler.GenerateAttendancesReportXlsx)
}
