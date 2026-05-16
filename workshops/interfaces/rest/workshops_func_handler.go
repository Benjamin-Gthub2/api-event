package rest

import (
	"errors"
	"net/http"

	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

// GetWorkshopById is a method to get workshop by id
// @Summary Get workshop by id
// @Description Get workshop by id
// @Tags Workshops
// @Accept json
// @Produce json
// @Param workshopId path string true "the id of the workshop"
// @Success 200 {object} workshopByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops/{workshopId} [get]
// @Security BearerAuth
func (h workshopsHandler) GetWorkshopById(c *gin.Context) {
	ctx := c.Request.Context()
	workshopId := c.Param("workshopId")

	workshop, err := h.workshopsUseCase.GetWorkshopById(ctx, workshopId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopByIdResult{
		Data:   workshop,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetWorkshops is a method to get workshops
// @Summary Get workshops
// @Description Get workshops
// @Tags Workshops
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param event_id query string false "Filter by event id"
// @Param type_id query string false "Filter by workshop type id"
// @Success 200 {object} workshopsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops [get]
// @Security BearerAuth
func (h workshopsHandler) GetWorkshops(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := workshopsDomain.GetWorkshopsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	workshops, paginationRes, err := h.workshopsUseCase.GetWorkshops(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopsResult{
		Data:       workshops,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateWorkshop is a method to create a workshop
// @Summary Create workshop
// @Description Create workshop
// @Tags Workshops
// @Accept json
// @Produce json
// @Param createWorkshopBody body createWorkshopValidated true "Create workshop body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops [post]
// @Security BearerAuth
func (h workshopsHandler) CreateWorkshop(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createWorkshopValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateWorkshop").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateWorkshop").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := workshopsDomain.CreateWorkshopBody{
		TypeId:    createBodyValidated.TypeId,
		Name:      createBodyValidated.Name,
		Shortname: createBodyValidated.Shortname,
		Code:      createBodyValidated.Code,
		Capacity:  createBodyValidated.Capacity,
		StartDate: createBodyValidated.StartDate,
		EndDate:   createBodyValidated.EndDate,
		Place:     createBodyValidated.Place,
		EventId:   createBodyValidated.EventId,
	}

	id, err := h.workshopsUseCase.CreateWorkshop(ctx, userId, createBody)
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

// UpdateWorkshop is a method to update a workshop
// @Summary Update workshop
// @Description Update workshop
// @Tags Workshops
// @Accept json
// @Produce json
// @Param workshopId path string true "the id of the workshop"
// @Param updateWorkshopBody body updateWorkshopValidated true "Update workshop body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops/{workshopId} [put]
// @Security BearerAuth
func (h workshopsHandler) UpdateWorkshop(c *gin.Context) {
	ctx := c.Request.Context()
	workshopId := c.Param("workshopId")

	var updateBodyValidated updateWorkshopValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateWorkshop").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateWorkshop").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := workshopsDomain.UpdateWorkshopBody{
		TypeId:    updateBodyValidated.TypeId,
		Name:      updateBodyValidated.Name,
		Shortname: updateBodyValidated.Shortname,
		Code:      updateBodyValidated.Code,
		Capacity:  updateBodyValidated.Capacity,
		StartDate: updateBodyValidated.StartDate,
		EndDate:   updateBodyValidated.EndDate,
		Place:     updateBodyValidated.Place,
	}

	err := h.workshopsUseCase.UpdateWorkshop(ctx, workshopId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   workshopId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteWorkshop is a method to delete a workshop
// @Summary Delete workshop
// @Description Delete workshop (soft delete)
// @Tags Workshops
// @Accept json
// @Produce json
// @Param workshopId path string true "the id of the workshop"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops/{workshopId} [delete]
// @Security BearerAuth
func (h workshopsHandler) DeleteWorkshop(c *gin.Context) {
	ctx := c.Request.Context()
	workshopId := c.Param("workshopId")
	userId := c.GetString("userId")

	err := h.workshopsUseCase.DeleteWorkshop(ctx, workshopId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   workshopId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetWorkshopsSummary is a method to get workshop sums
// @Summary Get workshop sums
// @Description Get workshop sums
// @Tags WorkshopSums
// @Accept json
// @Produce json
// @Param workshop_id query string false "the id of workshop"
// @Param search_value query string false "the search value of workshop"
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {object} workshopsSummaryResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshops/summary [get]
// @Security BearerAuth
func (h workshopsHandler) GetWorkshopsSummary(c *gin.Context) {
	ctx := c.Request.Context()
	searchParams := workshopsDomain.GetWorkshopSumsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	workshopsSummary, err := h.workshopsUseCase.GetWorkshopSummary(ctx, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopsSummaryResult{
		Data:   workshopsSummary,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
