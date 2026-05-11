package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	swaggerRest "github.com/Benjamin-Gthub2/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/workshop-speakers/docs"
	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

type workshopSpeakersHandler struct {
	workshopSpeakersUseCase workshopSpeakersDomain.WorkshopSpeakersUseCase
	authMiddleware          authMiddleware.AuthMiddleware
	err                     *errDomain.SmartError
}

func NewWorkshopSpeakersHandler(
	workshopSpeakers workshopSpeakersDomain.WorkshopSpeakersUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &workshopSpeakersHandler{
		workshopSpeakersUseCase: workshopSpeakers,
		authMiddleware:          authMiddleware,
		err:                     errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfoworkshopSpeakers, docs.DocTemplateJson, "event", "workshop-speakers")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/workshop-speakers", handler.GetWorkshopSpeakers)
	api.GET("/workshop-speakers/:workshopSpeakerId", handler.GetWorkshopSpeakerById)
	api.POST("/workshop-speakers", handler.CreateWorkshopSpeaker)
	api.DELETE("/workshop-speakers/:workshopSpeakerId", handler.DeleteWorkshopSpeaker)
}
