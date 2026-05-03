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
	"errors"
	"net/http"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	httpResponse "github.com/smart0n3/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"

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
// @Router /api/v1/event/registrations/{registrationId}/qr [get]
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

// GetRegistrationById is a method to get registration by id
// @Summary Get registration by id
// @Description Get registration by id
// @Tags Registrations
// @Accept json
// @Produce json
// @Param registrationId query string false "the id of the registration"
// @Success 200 {object} registrationByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registrations/{registrationId} [get]
// @Security BearerAuth
func (h registrationsHandler) GetRegistrationById(c *gin.Context) {
	ctx := c.Request.Context()
	registrationId := c.Param("registrationId")

	registrationById, err := h.registrationsUseCase.GetRegistrationById(ctx, registrationId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := registrationByIdResult{
		Data:   registrationById,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetRegistrations is a method to get registrations
// @Summary Get registrations
// @Description Get registrations
// @Tags Registrations
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param start_date query string false "the start date"
// @Param end_date query string false "the end date"
// @Param created_by query string false "the creator of the registration"
// @Success 200 {object} registrationsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registrations [get]
// @Security BearerAuth
func (h registrationsHandler) GetRegistrations(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := registrationsDomain.GetRegistrationsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	registrations, paginationRes, err := h.registrationsUseCase.GetRegistrations(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := registrationsResult{
		Data:       registrations,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateRegistration is a method to create registrations
// @Summary Create registrations
// @Description Create registrations
// @Tags Registrations
// @Accept json
// @Produce json
// @Param storeId path string true "Store id"
// @Param createRegistrationBody body registrationsDomain.CreateRegistrationBody true "Create registrations body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registrations [post]
// @Security BearerAuth
func (h registrationsHandler) CreateRegistration(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createRegistrationValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateRegistration").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateRegistration").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createBody = registrationsDomain.CreateRegistrationBody{
		SessionId:     createBodyValidated.SessionId,
		BeneficiaryId: createBodyValidated.BeneficiaryId,
	}

	id, err := h.registrationsUseCase.CreateRegistration(ctx, userId, createBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   *id,
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusCreated, res)
}

// UpdateRegistrationApprovalStatus is a method to update registration approval status
// @Summary Update registration approval status
// @Description Update registration approval status
// @Tags Registrations
// @Accept json
// @Produce json
// @Param registrationId path string true "registration id"
// @Param statusCode path string true "status code"
// @Success 201 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registrations/{registrationId}/statuses/{statusCode} [put]
// @Security BearerAuth
func (h registrationsHandler) UpdateRegistrationApprovalStatus(c *gin.Context) {
	ctx := c.Request.Context()
	registrationId := c.Param("registrationId")
	statusCode := c.Param("statusCode")
	err := h.registrationsUseCase.UpdateRegistrationStatus(ctx, registrationId, statusCode)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
