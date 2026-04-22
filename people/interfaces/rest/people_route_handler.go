/*
 * File: people_route_handler.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/people/docs"
	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

type peopleHandler struct {
	peopleUseCase  domain.PeopleUseCase
	authMiddleware authMiddleware.AuthMiddleware
	err            *errDomain.SmartError
}

func NewPeopleHandler(
	people domain.PeopleUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &peopleHandler{
		peopleUseCase:  people,
		authMiddleware: authMiddleware,
		err:            errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInfopeople, docs.DocTemplateJson, "event", "people")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)

	api.GET("/people/", handler.GetPeople)
	api.POST("/people/", handler.CreatePerson)
	api.PUT("/people/:personId", handler.UpdatePerson)
	api.DELETE("/people/:personId", handler.DeletePerson)
}
