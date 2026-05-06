package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	swaggerRest "github.com/Benjamin-Gthub2/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/registration-statuses/docs"
	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

type registrationStatusesHandler struct {
	registrationStatusesUseCase registrationStatusesDomain.RegistrationStatusesUseCase
	authMiddleware              authMiddleware.AuthMiddleware
	err                         *errDomain.SmartError
}

func NewRegistrationStatusesHandler(
	registrationStatuses registrationStatusesDomain.RegistrationStatusesUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &registrationStatusesHandler{
		registrationStatusesUseCase: registrationStatuses,
		authMiddleware:              authMiddleware,
		err:                         errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInforegistrationStatuses, docs.DocTemplateJson, "event", "registration-statuses")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/registration-statuses", handler.GetRegistrationStatuses)
	api.GET("/registration-statuses/:registrationStatusId", handler.GetRegistrationStatusById)
	api.POST("/registration-statuses", handler.CreateRegistrationStatus)
	api.PUT("/registration-statuses/:registrationStatusId", handler.UpdateRegistrationStatus)
	api.DELETE("/registration-statuses/:registrationStatusId", handler.DeleteRegistrationStatus)
}
