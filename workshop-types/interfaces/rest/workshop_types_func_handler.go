package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

// GetWorkshopTypeById is a method to get workshop type by id
// @Summary Get workshop type by id
// @Description Get workshop type by id
// @Tags WorkshopTypes
// @Accept json
// @Produce json
// @Param workshopTypeId path string true "the id of the workshop type"
// @Success 200 {object} workshopTypeByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-types/{workshopTypeId} [get]
// @Security BearerAuth
func (h workshopTypesHandler) GetWorkshopTypeById(c *gin.Context) {
	ctx := c.Request.Context()
	workshopTypeId := c.Param("workshopTypeId")

	workshopType, err := h.workshopTypesUseCase.GetWorkshopTypeById(ctx, workshopTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopTypeByIdResult{
		Data:   workshopType,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetWorkshopTypes is a method to get workshop types
// @Summary Get workshop types
// @Description Get workshop types
// @Tags WorkshopTypes
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Success 200 {object} workshopTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-types [get]
// @Security BearerAuth
func (h workshopTypesHandler) GetWorkshopTypes(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := workshopTypesDomain.GetWorkshopTypesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	workshopTypes, paginationRes, err := h.workshopTypesUseCase.GetWorkshopTypes(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopTypesResult{
		Data:       workshopTypes,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateWorkshopType is a method to create a workshop type
// @Summary Create workshop type
// @Description Create workshop type
// @Tags WorkshopTypes
// @Accept json
// @Produce json
// @Param createWorkshopTypeBody body createWorkshopTypeValidated true "Create workshop type body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-types [post]
// @Security BearerAuth
func (h workshopTypesHandler) CreateWorkshopType(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createWorkshopTypeValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateWorkshopType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateWorkshopType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := workshopTypesDomain.CreateWorkshopTypeBody{
		Code:        createBodyValidated.Code,
		Description: createBodyValidated.Description,
		Enable:      createBodyValidated.Enable,
	}

	id, err := h.workshopTypesUseCase.CreateWorkshopType(ctx, userId, createBody)
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

// UpdateWorkshopType is a method to update a workshop type
// @Summary Update workshop type
// @Description Update workshop type
// @Tags WorkshopTypes
// @Accept json
// @Produce json
// @Param workshopTypeId path string true "the id of the workshop type"
// @Param updateWorkshopTypeBody body updateWorkshopTypeValidated true "Update workshop type body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-types/{workshopTypeId} [put]
// @Security BearerAuth
func (h workshopTypesHandler) UpdateWorkshopType(c *gin.Context) {
	ctx := c.Request.Context()
	workshopTypeId := c.Param("workshopTypeId")

	var updateBodyValidated updateWorkshopTypeValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateWorkshopType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateWorkshopType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := workshopTypesDomain.UpdateWorkshopTypeBody{
		Code:        updateBodyValidated.Code,
		Description: updateBodyValidated.Description,
		Enable:      updateBodyValidated.Enable,
	}

	err := h.workshopTypesUseCase.UpdateWorkshopType(ctx, workshopTypeId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   workshopTypeId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteWorkshopType is a method to delete a workshop type
// @Summary Delete workshop type
// @Description Delete workshop type (soft delete)
// @Tags WorkshopTypes
// @Accept json
// @Produce json
// @Param workshopTypeId path string true "the id of the workshop type"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-types/{workshopTypeId} [delete]
// @Security BearerAuth
func (h workshopTypesHandler) DeleteWorkshopType(c *gin.Context) {
	ctx := c.Request.Context()
	workshopTypeId := c.Param("workshopTypeId")
	userId := c.GetString("userId")

	err := h.workshopTypesUseCase.DeleteWorkshopType(ctx, workshopTypeId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   workshopTypeId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
