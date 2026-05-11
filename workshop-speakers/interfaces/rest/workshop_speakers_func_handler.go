package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

// GetWorkshopSpeakerById is a method to get workshop speaker by id
// @Summary Get workshop speaker by id
// @Description Get workshop speaker by id
// @Tags WorkshopSpeakers
// @Accept json
// @Produce json
// @Param workshopSpeakerId path string true "the id of the workshop speaker"
// @Success 200 {object} workshopSpeakerByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-speakers/{workshopSpeakerId} [get]
// @Security BearerAuth
func (h workshopSpeakersHandler) GetWorkshopSpeakerById(c *gin.Context) {
	ctx := c.Request.Context()
	workshopSpeakerId := c.Param("workshopSpeakerId")

	workshopSpeaker, err := h.workshopSpeakersUseCase.GetWorkshopSpeakerById(ctx, workshopSpeakerId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopSpeakerByIdResult{
		Data:   workshopSpeaker,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetWorkshopSpeakers is a method to get workshop speakers
// @Summary Get workshop speakers
// @Description Get workshop speakers
// @Tags WorkshopSpeakers
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param workshop_id query string false "Filter by workshop id"
// @Param speaker_id query string false "Filter by speaker id"
// @Success 200 {object} workshopSpeakersResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-speakers [get]
// @Security BearerAuth
func (h workshopSpeakersHandler) GetWorkshopSpeakers(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := workshopSpeakersDomain.GetWorkshopSpeakersParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	workshopSpeakers, paginationRes, err := h.workshopSpeakersUseCase.GetWorkshopSpeakers(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := workshopSpeakersResult{
		Data:       workshopSpeakers,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateWorkshopSpeaker is a method to create a workshop speaker
// @Summary Create workshop speaker
// @Description Create workshop speaker
// @Tags WorkshopSpeakers
// @Accept json
// @Produce json
// @Param createWorkshopSpeakerBody body createWorkshopSpeakerValidated true "Create workshop speaker body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-speakers [post]
// @Security BearerAuth
func (h workshopSpeakersHandler) CreateWorkshopSpeaker(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createWorkshopSpeakerValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateWorkshopSpeaker").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateWorkshopSpeaker").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := workshopSpeakersDomain.CreateWorkshopSpeakerBody{
		WorkshopId:         createBodyValidated.WorkshopId,
		SpeakerId:          createBodyValidated.SpeakerId,
		DegreeAbbreviation: createBodyValidated.DegreeAbbreviation,
	}

	id, err := h.workshopSpeakersUseCase.CreateWorkshopSpeaker(ctx, userId, createBody)
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

// DeleteWorkshopSpeaker is a method to delete a workshop speaker
// @Summary Delete workshop speaker
// @Description Delete workshop speaker (soft delete)
// @Tags WorkshopSpeakers
// @Accept json
// @Produce json
// @Param workshopSpeakerId path string true "the id of the workshop speaker"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/workshop-speakers/{workshopSpeakerId} [delete]
// @Security BearerAuth
func (h workshopSpeakersHandler) DeleteWorkshopSpeaker(c *gin.Context) {
	ctx := c.Request.Context()
	workshopSpeakerId := c.Param("workshopSpeakerId")
	userId := c.GetString("userId")

	err := h.workshopSpeakersUseCase.DeleteWorkshopSpeaker(ctx, workshopSpeakerId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   workshopSpeakerId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
