package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

// GetMaterialIssuedById is a method to get material issued by id
// @Summary Get material issued by id
// @Description Get material issued by id
// @Tags MaterialsIssued
// @Accept json
// @Produce json
// @Param materialIssuedId path string true "the id of the material issued"
// @Success 200 {object} materialIssuedByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/materials-issued/{materialIssuedId} [get]
// @Security BearerAuth
func (h materialsIssuedHandler) GetMaterialIssuedById(c *gin.Context) {
	ctx := c.Request.Context()
	materialIssuedId := c.Param("materialIssuedId")

	materialIssued, err := h.materialsIssuedUseCase.GetMaterialIssuedById(ctx, materialIssuedId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := materialIssuedByIdResult{
		Data:   materialIssued,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetMaterialsIssued is a method to get materials issued
// @Summary Get materials issued
// @Description Get materials issued
// @Tags MaterialsIssued
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param start_date query string false "Start date filter"
// @Param end_date query string false "End date filter"
// @Success 200 {object} materialsIssuedResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/materials-issued [get]
// @Security BearerAuth
func (h materialsIssuedHandler) GetMaterialsIssued(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := materialsIssuedDomain.GetMaterialsIssuedParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	materialsIssued, paginationRes, err := h.materialsIssuedUseCase.GetMaterialsIssued(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := materialsIssuedResult{
		Data:       materialsIssued,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateMaterialIssued is a method to create a material issued
// @Summary Create material issued
// @Description Create material issued
// @Tags MaterialsIssued
// @Accept json
// @Produce json
// @Param createMaterialIssuedBody body createMaterialIssuedValidated false "Create material issued body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/materials-issued [post]
// @Security BearerAuth
func (h materialsIssuedHandler) CreateMaterialIssued(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createMaterialIssuedValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateMaterialIssued").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateMaterialIssued").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := materialsIssuedDomain.CreateMaterialIssuedBody{
		Description: createBodyValidated.Description,
	}

	id, err := h.materialsIssuedUseCase.CreateMaterialIssued(ctx, userId, createBody)
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

// UpdateMaterialIssued is a method to update a material issued
// @Summary Update material issued
// @Description Update material issued
// @Tags MaterialsIssued
// @Accept json
// @Produce json
// @Param materialIssuedId path string true "the id of the material issued"
// @Param updateMaterialIssuedBody body updateMaterialIssuedValidated false "Update material issued body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/materials-issued/{materialIssuedId} [put]
// @Security BearerAuth
func (h materialsIssuedHandler) UpdateMaterialIssued(c *gin.Context) {
	ctx := c.Request.Context()
	materialIssuedId := c.Param("materialIssuedId")

	var updateBodyValidated updateMaterialIssuedValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateMaterialIssued").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateMaterialIssued").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := materialsIssuedDomain.UpdateMaterialIssuedBody{
		Description: updateBodyValidated.Description,
	}

	err := h.materialsIssuedUseCase.UpdateMaterialIssued(ctx, materialIssuedId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   materialIssuedId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteMaterialIssued is a method to delete a material issued
// @Summary Delete material issued
// @Description Delete material issued (soft delete)
// @Tags MaterialsIssued
// @Accept json
// @Produce json
// @Param materialIssuedId path string true "the id of the material issued"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/materials-issued/{materialIssuedId} [delete]
// @Security BearerAuth
func (h materialsIssuedHandler) DeleteMaterialIssued(c *gin.Context) {
	ctx := c.Request.Context()
	materialIssuedId := c.Param("materialIssuedId")
	userId := c.GetString("userId")

	err := h.materialsIssuedUseCase.DeleteMaterialIssued(ctx, materialIssuedId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   materialIssuedId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
