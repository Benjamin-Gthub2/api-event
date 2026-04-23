package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/smart0n3/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/smart0n3/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

// GetEventTypeById is a method to get event type by id
// @Summary Get event type by id
// @Description Get event type by id
// @Tags EventTypes
// @Accept json
// @Produce json
// @Param eventTypeId path string true "the id of the event type"
// @Success 200 {object} eventTypeByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/event-types/{eventTypeId} [get]
// @Security BearerAuth
func (h eventTypesHandler) GetEventTypeById(c *gin.Context) {
	ctx := c.Request.Context()
	eventTypeId := c.Param("eventTypeId")

	eventType, err := h.eventTypesUseCase.GetEventTypeById(ctx, eventTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := eventTypeByIdResult{
		Data:   eventType,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetEventTypes is a method to get event types
// @Summary Get event types
// @Description Get event types
// @Tags EventTypes
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Success 200 {object} eventTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/event-types [get]
// @Security BearerAuth
func (h eventTypesHandler) GetEventTypes(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := eventTypesDomain.GetEventTypesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	eventTypes, paginationRes, err := h.eventTypesUseCase.GetEventTypes(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := eventTypesResult{
		Data:       eventTypes,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateEventType is a method to create an event type
// @Summary Create event type
// @Description Create event type
// @Tags EventTypes
// @Accept json
// @Produce json
// @Param createEventTypeBody body createEventTypeValidated true "Create event type body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/event-types [post]
// @Security BearerAuth
func (h eventTypesHandler) CreateEventType(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createEventTypeValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateEventType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateEventType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := eventTypesDomain.CreateEventTypeBody{
		Code:        createBodyValidated.Code,
		Description: createBodyValidated.Description,
		Enable:      createBodyValidated.Enable,
	}

	id, err := h.eventTypesUseCase.CreateEventType(ctx, userId, createBody)
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

// UpdateEventType is a method to update an event type
// @Summary Update event type
// @Description Update event type
// @Tags EventTypes
// @Accept json
// @Produce json
// @Param eventTypeId path string true "the id of the event type"
// @Param updateEventTypeBody body updateEventTypeValidated true "Update event type body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/event-types/{eventTypeId} [put]
// @Security BearerAuth
func (h eventTypesHandler) UpdateEventType(c *gin.Context) {
	ctx := c.Request.Context()
	eventTypeId := c.Param("eventTypeId")

	var updateBodyValidated updateEventTypeValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateEventType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateEventType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := eventTypesDomain.UpdateEventTypeBody{
		Code:        updateBodyValidated.Code,
		Description: updateBodyValidated.Description,
		Enable:      updateBodyValidated.Enable,
	}

	err := h.eventTypesUseCase.UpdateEventType(ctx, eventTypeId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   eventTypeId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteEventType is a method to delete an event type
// @Summary Delete event type
// @Description Delete event type (soft delete)
// @Tags EventTypes
// @Accept json
// @Produce json
// @Param eventTypeId path string true "the id of the event type"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/event-types/{eventTypeId} [delete]
// @Security BearerAuth
func (h eventTypesHandler) DeleteEventType(c *gin.Context) {
	ctx := c.Request.Context()
	eventTypeId := c.Param("eventTypeId")
	userId := c.GetString("userId")

	err := h.eventTypesUseCase.DeleteEventType(ctx, eventTypeId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   eventTypeId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
