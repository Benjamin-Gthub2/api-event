package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	swaggerRest "github.com/Benjamin-Gthub2/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/workshop-types/docs"
	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

type workshopTypesHandler struct {
	workshopTypesUseCase workshopTypesDomain.WorkshopTypesUseCase
	authMiddleware       authMiddleware.AuthMiddleware
	err                  *errDomain.SmartError
}

func NewWorkshopTypesHandler(
	workshopTypes workshopTypesDomain.WorkshopTypesUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &workshopTypesHandler{
		workshopTypesUseCase: workshopTypes,
		authMiddleware:       authMiddleware,
		err:                  errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfoworkshop_types, docs.DocTemplateJson, "event", "workshop-types")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/workshop-types", handler.GetWorkshopTypes)
	api.GET("/workshop-types/:workshopTypeId", handler.GetWorkshopTypeById)
	api.POST("/workshop-types", handler.CreateWorkshopType)
	api.PUT("/workshop-types/:workshopTypeId", handler.UpdateWorkshopType)
	api.DELETE("/workshop-types/:workshopTypeId", handler.DeleteWorkshopType)
}
