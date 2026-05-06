package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

// GetRegistrationStatusById is a method to get registration status by id
// @Summary Get registration status by id
// @Description Get registration status by id
// @Tags RegistrationStatuses
// @Accept json
// @Produce json
// @Param registrationStatusId path string true "the id of the registration status"
// @Success 200 {object} registrationStatusByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registration-statuses/{registrationStatusId} [get]
// @Security BearerAuth
func (h registrationStatusesHandler) GetRegistrationStatusById(c *gin.Context) {
	ctx := c.Request.Context()
	registrationStatusId := c.Param("registrationStatusId")

	registrationStatus, err := h.registrationStatusesUseCase.GetRegistrationStatusById(ctx, registrationStatusId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := registrationStatusByIdResult{
		Data:   registrationStatus,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetRegistrationStatuses is a method to get registration statuses
// @Summary Get registration statuses
// @Description Get registration statuses
// @Tags RegistrationStatuses
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Success 200 {object} registrationStatusesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registration-statuses [get]
// @Security BearerAuth
func (h registrationStatusesHandler) GetRegistrationStatuses(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := registrationStatusesDomain.GetRegistrationStatusesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	registrationStatuses, paginationRes, err := h.registrationStatusesUseCase.GetRegistrationStatuses(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := registrationStatusesResult{
		Data:       registrationStatuses,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateRegistrationStatus is a method to create a registration status
// @Summary Create registration status
// @Description Create registration status
// @Tags RegistrationStatuses
// @Accept json
// @Produce json
// @Param createRegistrationStatusBody body createRegistrationStatusValidated true "Create registration status body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registration-statuses [post]
// @Security BearerAuth
func (h registrationStatusesHandler) CreateRegistrationStatus(c *gin.Context) {
	ctx := c.Request.Context()

	var createBodyValidated createRegistrationStatusValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateRegistrationStatus").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateRegistrationStatus").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := registrationStatusesDomain.CreateRegistrationStatusBody{
		Code:        createBodyValidated.Code,
		Description: createBodyValidated.Description,
		Position:    createBodyValidated.Position,
		Enable:      createBodyValidated.Enable,
	}

	id, err := h.registrationStatusesUseCase.CreateRegistrationStatus(ctx, createBody)
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

// UpdateRegistrationStatus is a method to update a registration status
// @Summary Update registration status
// @Description Update registration status
// @Tags RegistrationStatuses
// @Accept json
// @Produce json
// @Param registrationStatusId path string true "the id of the registration status"
// @Param updateRegistrationStatusBody body updateRegistrationStatusValidated true "Update registration status body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registration-statuses/{registrationStatusId} [put]
// @Security BearerAuth
func (h registrationStatusesHandler) UpdateRegistrationStatus(c *gin.Context) {
	ctx := c.Request.Context()
	registrationStatusId := c.Param("registrationStatusId")

	var updateBodyValidated updateRegistrationStatusValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateRegistrationStatus").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateRegistrationStatus").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := registrationStatusesDomain.UpdateRegistrationStatusBody{
		Code:        updateBodyValidated.Code,
		Description: updateBodyValidated.Description,
		Position:    updateBodyValidated.Position,
		Enable:      updateBodyValidated.Enable,
	}

	err := h.registrationStatusesUseCase.UpdateRegistrationStatus(ctx, registrationStatusId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   registrationStatusId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteRegistrationStatus is a method to delete a registration status
// @Summary Delete registration status
// @Description Delete registration status (soft delete)
// @Tags RegistrationStatuses
// @Accept json
// @Produce json
// @Param registrationStatusId path string true "the id of the registration status"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/registration-statuses/{registrationStatusId} [delete]
// @Security BearerAuth
func (h registrationStatusesHandler) DeleteRegistrationStatus(c *gin.Context) {
	ctx := c.Request.Context()
	registrationStatusId := c.Param("registrationStatusId")

	err := h.registrationStatusesUseCase.DeleteRegistrationStatus(ctx, registrationStatusId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   registrationStatusId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
