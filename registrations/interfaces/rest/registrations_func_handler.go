/*
 * File: registrations_receipts_func_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the func handler of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	restCore "github.com/smart0n3/api-shared/api-core/interfaces/rest"
)

// GetQrRegistrationById is a method to get qr registrations
// @Summary Get qr registration by id
// @Description Get qr registration by id
// @Tags Registrations
// @Accept json
// @Produce image/png
// @Param registrationId query string false "the id of the registration"
// @Success 200 {string} string "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registrations/{registrationId} [get]
// @Security BearerAuth
func (h registrationsHandler) GetQrRegistrationById(c *gin.Context) {
	ctx := c.Request.Context()
	registrationId := c.Param("registrationId")

	qrRegistration, err := h.registrationsUseCase.GetQrRegistrationById(ctx, registrationId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	c.Data(http.StatusOK, "image/png", qrRegistration)
}
