package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	swaggerRest "github.com/Benjamin-Gthub2/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/attendances/docs"
	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendancesHandler struct {
	attendancesUseCase attendancesDomain.AttendancesUseCase
	authMiddleware     authMiddleware.AuthMiddleware
	err                *errDomain.SmartError
}

func NewAttendancesHandler(
	attendances attendancesDomain.AttendancesUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &attendancesHandler{
		attendancesUseCase: attendances,
		authMiddleware:     authMiddleware,
		err:                errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfoattendances, docs.DocTemplateJson, "event", "attendances")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/attendances", handler.GetAttendances)
	api.GET("/attendances/:attendanceId", handler.GetAttendanceById)
	api.POST("/attendances", handler.CreateAttendance)
	api.DELETE("/attendances/:attendanceId", handler.DeleteAttendance)
}
