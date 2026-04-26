/*
 * File: events_func_handler.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains handler functions for managing events related operations.
 *
 * Last Modified: 2026-04-15
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "github.com/smart0n3/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/smart0n3/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"

	eventsDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

// GetEvents is a method to get event
// @Summary Get events
// @Description Get event
// @Tags Events
// @Accept json
// @Produce json
// @Param status query string false "the status of the event"
// @Param name_or_document query string false "the name or document of the event"
// @Success 200 {object} eventsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events [get]
// @Security BearerAuth
func (h eventsHandler) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := eventsDomain.GetEventsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	events, paginationRes, err := h.eventsUseCase.GetEvents(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := eventsResult{
		Data:       events,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateEvent is a method to create event
// @Summary Create event
// @Description Create event
// @Tags Events
// @Accept json
// @Produce json
// @Param createEventBody body eventsDomain.CreateEventBody true "Create event body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events [post]
// @Security BearerAuth
func (h eventsHandler) CreateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var eventsValidate createEventsValidate
	if err := c.ShouldBindJSON(&eventsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateEvent").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateEvent").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var createEventBody = eventsDomain.CreateEventBody{
		Name:        eventsValidate.Name,
		Description: eventsValidate.Description,
		Code:        eventsValidate.Code,
		Phone:       eventsValidate.Phone,
		Document:    eventsValidate.Document,
		Address:     eventsValidate.Address,
		Industry:    eventsValidate.Industry,
		Enable:      eventsValidate.Enable,
	}
	id, err := h.eventsUseCase.CreateEvent(ctx, createEventBody)
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

// UpdateEvent is a method to update event
// @Summary Update event
// @Description Update event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventId path string true "event id"
// @Param updateEventBody body eventsDomain.UpdateEventBody true "Update event body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events/{eventId} [put]
// @Security BearerAuth
func (h eventsHandler) UpdateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	eventId := c.Param("eventId")

	var eventsValidate createEventsValidate
	if err := c.ShouldBindJSON(&eventsValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateEvent").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateEvent").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var eventsBody = eventsDomain.UpdateEventBody{
		Name:        eventsValidate.Name,
		Description: eventsValidate.Description,
		Code:        eventsValidate.Code,
		Phone:       eventsValidate.Phone,
		Document:    eventsValidate.Document,
		Address:     eventsValidate.Address,
		Industry:    eventsValidate.Industry,
		Enable:      eventsValidate.Enable,
	}
	err := h.eventsUseCase.UpdateEvent(ctx, eventId, eventsBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteEvent is a method to delete event
// @Summary Delete a event
// @Description Delete event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventId path string true "event id"
// @Success 200 {object} deleteEventsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events/{eventId} [delete]
// @Security BearerAuth
func (h eventsHandler) DeleteEvent(c *gin.Context) {
	ctx := c.Request.Context()
	eventId := c.Param("eventId")
	result, err := h.eventsUseCase.DeleteEvent(ctx, eventId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteEventsResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetRolesByEvent is a method to get roles by event
// @Summary Get roles by event
// @Description Get roles by event
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} rolesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events/{eventId}/roles [get]
// @Security BearerAuth
func (h eventsHandler) GetRolesByEvent(c *gin.Context) {
	ctx := c.Request.Context()
	eventId := c.Param("eventId")

	roles, err := h.eventsUseCase.GetRolesByEvent(ctx, eventId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := rolesResult{
		Data:   roles,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// ToggleEventEnable is a method to enable or disable a event
// @Summary Enable or disable event
// @Description Enable or disable event
// @Tags Events
// @Accept json
// @Produce json
// @Param eventId path string true "event id"
// @Param enableDisableEventBody body eventsDomain.EnableDisableEventRequest true "Enable or Disable event body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/events/{eventId}/enable [put]
// @Security BearerAuth
func (h eventsHandler) ToggleEventEnable(c *gin.Context) {
	ctx := c.Request.Context()
	eventId := c.Param("eventId")

	var enableValidate enableDisableEventValidate
	if err := c.ShouldBindJSON(&enableValidate); err != nil {
		validationErrs, isValidationErr := err.(validator.ValidationErrors)
		if !isValidationErr {
			err = h.err.Clone().SetFunction("EnableDisableEvent").SetRaw(err)
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("EnableDisableEvent").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	var body = eventsDomain.EnableDisableEventRequest{
		Enable: enableValidate.Enable,
	}

	err := h.eventsUseCase.EnableDisableEvent(ctx, eventId, body)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetEventsSummary is a method to get event sums
// @Summary Get event sums
// @Description Get event sums
// @Tags EventSums
// @Accept json
// @Produce json
// @Param event_id query string false "the id of event"
// @Success 200 {object} eventsSummaryResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/events/summary [get]
// @Security BearerAuth
func (h eventsHandler) GetEventsSummary(c *gin.Context) {
	ctx := c.Request.Context()
	searchParams := eventsDomain.GetEventSumsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	eventsSummary, err := h.eventsUseCase.GetEventSummary(ctx, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := eventsSummaryResult{
		Data:   eventsSummary,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
