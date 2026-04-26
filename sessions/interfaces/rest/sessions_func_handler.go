package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	restCore "github.com/smart0n3/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/smart0n3/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

// GetSessionById is a method to get session by id
// @Summary Get session by id
// @Description Get session by id
// @Tags Sessions
// @Accept json
// @Produce json
// @Param sessionId path string true "the id of the session"
// @Success 200 {object} sessionByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions/{sessionId} [get]
// @Security BearerAuth
func (h sessionsHandler) GetSessionById(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := c.Param("sessionId")

	session, err := h.sessionsUseCase.GetSessionById(ctx, sessionId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := sessionByIdResult{
		Data:   session,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetSessions is a method to get sessions
// @Summary Get sessions
// @Description Get sessions
// @Tags Sessions
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param workshop_id query string false "Filter by workshop id"
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {object} sessionsResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions [get]
// @Security BearerAuth
func (h sessionsHandler) GetSessions(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := sessionsDomain.GetSessionsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	sessions, paginationRes, err := h.sessionsUseCase.GetSessions(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := sessionsResult{
		Data:       sessions,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateSession is a method to create a session
// @Summary Create session
// @Description Create session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param createSessionBody body createSessionValidated true "Create session body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions [post]
// @Security BearerAuth
func (h sessionsHandler) CreateSession(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	var createBodyValidated createSessionValidated
	if err := c.ShouldBindJSON(&createBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateSession").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateSession").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createBody := sessionsDomain.CreateSessionBody{
		WorkshopId: createBodyValidated.WorkshopId,
		StartDate:  createBodyValidated.StartDate,
		EndDate:    createBodyValidated.EndDate,
	}

	id, err := h.sessionsUseCase.CreateSession(ctx, userId, createBody)
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

// UpdateSession is a method to update a session
// @Summary Update session
// @Description Update session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param sessionId path string true "the id of the session"
// @Param updateSessionBody body updateSessionValidated true "Update session body"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions/{sessionId} [put]
// @Security BearerAuth
func (h sessionsHandler) UpdateSession(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := c.Param("sessionId")

	var updateBodyValidated updateSessionValidated
	if err := c.ShouldBindJSON(&updateBodyValidated); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateSession").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}
		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateSession").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	updateBody := sessionsDomain.UpdateSessionBody{
		StartDate: updateBodyValidated.StartDate,
		EndDate:   updateBodyValidated.EndDate,
	}

	err := h.sessionsUseCase.UpdateSession(ctx, sessionId, updateBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   sessionId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteSession is a method to delete a session
// @Summary Delete session
// @Description Delete session (soft delete)
// @Tags Sessions
// @Accept json
// @Produce json
// @Param sessionId path string true "the id of the session"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions/{sessionId} [delete]
// @Security BearerAuth
func (h sessionsHandler) DeleteSession(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := c.Param("sessionId")
	userId := c.GetString("userId")

	err := h.sessionsUseCase.DeleteSession(ctx, sessionId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   sessionId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetSessionsSummary is a method to get session sums
// @Summary Get session sums
// @Description Get session sums
// @Tags SessionSums
// @Accept json
// @Produce json
// @Param session_id query string false "the id of session"
// @Success 200 {object} sessionsSummaryResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/sessions/summary [get]
// @Security BearerAuth
func (h sessionsHandler) GetSessionsSummary(c *gin.Context) {
	ctx := c.Request.Context()
	searchParams := sessionsDomain.GetSessionSumsParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	sessionsSummary, err := h.sessionsUseCase.GetSessionSummary(ctx, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := sessionsSummaryResult{
		Data:   sessionsSummary,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
