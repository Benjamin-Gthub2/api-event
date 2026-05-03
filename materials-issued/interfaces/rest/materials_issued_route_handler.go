package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/materials-issued/docs"
	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

type materialsIssuedHandler struct {
	materialsIssuedUseCase materialsIssuedDomain.MaterialsIssuedUseCase
	authMiddleware         authMiddleware.AuthMiddleware
	err                    *errDomain.SmartError
}

func NewMaterialsIssuedHandler(
	materialsIssued materialsIssuedDomain.MaterialsIssuedUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &materialsIssuedHandler{
		materialsIssuedUseCase: materialsIssued,
		authMiddleware:         authMiddleware,
		err:                    errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfomaterialsIssued, docs.DocTemplateJson, "event", "materials-issued")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/materials-issued", handler.GetMaterialsIssued)
	api.GET("/materials-issued/:materialIssuedId", handler.GetMaterialIssuedById)
	api.POST("/materials-issued", handler.CreateMaterialIssued)
	api.PUT("/materials-issued/:materialIssuedId", handler.UpdateMaterialIssued)
	api.DELETE("/materials-issued/:materialIssuedId", handler.DeleteMaterialIssued)
}
