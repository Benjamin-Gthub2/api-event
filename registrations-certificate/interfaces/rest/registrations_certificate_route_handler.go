/*
 * File: registrations_certificate_route_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the route handler for registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package rest

import (
	"github.com/gin-gonic/gin"

	authMiddleware "github.com/Benjamin-Gthub2/api-shared/auth/interfaces/rest"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationCertificateDomain "github.com/Benjamin-Gthub2/api-event/registrations-certificate/domain"
)

type registrationsCertificateHandler struct {
	registrationsCertificateUseCase registrationCertificateDomain.RegistrationsCertificateUseCase
	authMiddleware                  authMiddleware.AuthMiddleware
	err                             *errDomain.SmartError
}

func NewRegistrationsCertificateHandler(
	registrationsCertificateUseCase registrationCertificateDomain.RegistrationsCertificateUseCase,
	router *gin.Engine,
	authMiddleware authMiddleware.AuthMiddleware,
) {
	handler := &registrationsCertificateHandler{
		registrationsCertificateUseCase: registrationsCertificateUseCase,
		authMiddleware:                  authMiddleware,
		err:                             errDomain.NewErr().SetLayer(errDomain.Interface),
	}

	api := router.Group("/api/v1/event")
	api.Use(handler.authMiddleware.Auth)
	api.GET("/registrations/:registrationId/certificate", handler.GetCertificateByRegistrationId)
}
