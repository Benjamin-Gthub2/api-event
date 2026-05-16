/*
 * File: registrations_certificate_func_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the handler functions for registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
)

// GetCertificateByRegistrationId generates and downloads the certificate PDF for a registration
// @Summary Get certificate PDF by registration id
// @Description Generates and returns the participation certificate PDF for the given registration
// @Tags Registrations Certificate
// @Accept json
// @Produce application/pdf
// @Param registrationId path string true "the id of the registration"
// @Success 200 {file} binary "PDF certificate file"
// @Failure 400 {object} errorDomain.SmartError "Bad Request"
// @Failure 500 {object} errorDomain.SmartError "Internal Server Error"
// @Router /api/v1/event/registrations/{registrationId}/certificate [get]
// @Security BearerAuth
func (h registrationsCertificateHandler) GetCertificateByRegistrationId(c *gin.Context) {
	ctx := c.Request.Context()
	registrationId := c.Param("registrationId")

	pdfBytes, fileName, err := h.registrationsCertificateUseCase.GenerateRegistrationsCertificatePdf(ctx, registrationId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Header("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
