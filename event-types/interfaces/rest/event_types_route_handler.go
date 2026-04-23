package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/event-types/docs"
	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

type eventTypesHandler struct {
	eventTypesUseCase eventTypesDomain.EventTypesUseCase
	authMiddleware    authMiddleware.AuthMiddleware
	err               *errDomain.SmartError
}

func NewEventTypesHandler(
	eventTypes eventTypesDomain.EventTypesUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &eventTypesHandler{
		eventTypesUseCase: eventTypes,
		authMiddleware:    authMiddleware,
		err:               errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfoevent_types, docs.DocTemplateJson, "event", "event-types")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/event-types", handler.GetEventTypes)
	api.GET("/event-types/:eventTypeId", handler.GetEventTypeById)
	api.POST("/event-types", handler.CreateEventType)
	api.PUT("/event-types/:eventTypeId", handler.UpdateEventType)
	api.DELETE("/event-types/:eventTypeId", handler.DeleteEventType)
}
