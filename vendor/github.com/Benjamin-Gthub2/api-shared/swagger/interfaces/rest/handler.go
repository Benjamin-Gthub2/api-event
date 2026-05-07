package rest

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func Handler(router *gin.Engine, spec *swag.Spec, swaggerJson, service, microservice string) {
	title := fmt.Sprintf("Swagger %s - %s API", service, microservice)
	spec.Title = title
	description := fmt.Sprintf("This is a %s microservice.", microservice)
	spec.Description = description
	spec.Version = "1.0"
	spec.Host = "localhost:" + os.Getenv("SERVER_PORT")
	spec.BasePath = "/"
	spec.Schemes = []string{"http", "https"}
	instanceName := fmt.Sprintf("swagger_%s_%s", service, microservice)
	spec.InfoInstanceName = instanceName

	docV3Json := fmt.Sprintf("/%s/%s/docs/swagger3.json", service, microservice)
	router.GET(docV3Json, func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json", []byte(swaggerJson))
	})

	microserviceUnderScore := strings.ReplaceAll(microservice, "-", "_")
	swaggerRoot := fmt.Sprintf("/%s/%s/swagger/*any", service, microservice)
	router.GET(swaggerRoot, ginSwagger.WrapHandler(
		swaggerFiles.NewHandler(),
		ginSwagger.InstanceName(microserviceUnderScore),
		ginSwagger.URL(docV3Json),
	))
}
