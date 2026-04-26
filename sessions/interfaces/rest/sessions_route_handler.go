package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/sessions/docs"
	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

type sessionsHandler struct {
	sessionsUseCase sessionsDomain.SessionsUseCase
	authMiddleware  authMiddleware.AuthMiddleware
	err             *errDomain.SmartError
}

func NewSessionsHandler(
	sessions sessionsDomain.SessionsUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &sessionsHandler{
		sessionsUseCase: sessions,
		authMiddleware:  authMiddleware,
		err:             errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfosessions, docs.DocTemplateJson, "event", "sessions")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/sessions", handler.GetSessions)
	api.GET("/sessions/:sessionId", handler.GetSessionById)
	api.POST("/sessions", handler.CreateSession)
	api.PUT("/sessions/:sessionId", handler.UpdateSession)
	api.DELETE("/sessions/:sessionId", handler.DeleteSession)
	api.GET("/sessions/summary", handler.GetSessionsSummary)
}
