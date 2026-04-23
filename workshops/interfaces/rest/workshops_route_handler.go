package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/workshops/docs"
	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

type workshopsHandler struct {
	workshopsUseCase workshopsDomain.WorkshopsUseCase
	authMiddleware   authMiddleware.AuthMiddleware
	err              *errDomain.SmartError
}

func NewWorkshopsHandler(
	workshops workshopsDomain.WorkshopsUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &workshopsHandler{
		workshopsUseCase: workshops,
		authMiddleware:   authMiddleware,
		err:              errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfoworkshops, docs.DocTemplateJson, "event", "workshops")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/workshops", handler.GetWorkshops)
	api.GET("/workshops/:workshopId", handler.GetWorkshopById)
	api.POST("/workshops", handler.CreateWorkshop)
	api.PUT("/workshops/:workshopId", handler.UpdateWorkshop)
	api.DELETE("/workshops/:workshopId", handler.DeleteWorkshop)
	api.GET("/workshops/summary", handler.GetWorkshopsSummary)
}
