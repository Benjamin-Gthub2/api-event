/*
 * File: registrations_route_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the route handler of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/smart0n3/api-shared/auth/interfaces/rest"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	swaggerRest "github.com/smart0n3/api-shared/swagger/interfaces/rest"

	"github.com/Benjamin-Gthub2/api-event/registrations/docs"
	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsHandler struct {
	registrationsUseCase registrationsDomain.RegistrationsUseCase
	authMiddleware       authMiddleware.AuthMiddleware
	err                  *errDomain.SmartError
}

func NewRegistrationsHandler(
	registrations registrationsDomain.RegistrationsUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &registrationsHandler{
		registrationsUseCase: registrations,
		authMiddleware:       authMiddleware,
		err:                  errDomain.NewErr().SetLayer(errDomain.Interface),
	}
	swaggerRest.Handler(router, docs.SwaggerInforegistrations, docs.DocTemplateJson, "event", "registrations")

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/registrations", handler.GetRegistrations)
	api.GET("/registrations/:registrationId", handler.GetRegistrationById)
	api.GET("/registrations/:registrationId/qr", handler.GetQrRegistrationById) //mas adelante convertir en dos enpoints para get y post
	api.POST("/registrations", handler.CreateRegistration)
	api.PUT("/registrations/:registrationId/statuses/:statusCode", handler.UpdateRegistrationApprovalStatus)
}
