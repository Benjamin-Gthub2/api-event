/*
 * File: events_route_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the handler layer where the routes are located.
 *
 * Last Modified: 2026-04-15
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	swaggerRest "github.com/Benjamin-Gthub2/api-shared/swagger/interfaces/rest"

	eventsDocs "github.com/Benjamin-Gthub2/api-event/events/docs"
	eventsDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

type eventsHandler struct {
	eventsUseCase  eventsDomain.EventUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewEventsHandler(
	events eventsDomain.EventUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &eventsHandler{
		eventsUseCase:  events,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	swaggerRest.Handler(router, eventsDocs.SwaggerInfoevents, eventsDocs.DocTemplateJson,
		"event", "events")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/events", handler.GetEvents)
	api.POST("/events", handler.CreateEvent)
	api.PUT("/events/:eventId", handler.UpdateEvent)
	api.DELETE("/events/:eventId", handler.DeleteEvent)
	api.GET("/events/:eventId/roles", handler.GetRolesByEvent)
	api.PUT("/events/:eventId/enable", handler.ToggleEventEnable)
	api.GET("/events/summary", handler.GetEventsSummary)

}
