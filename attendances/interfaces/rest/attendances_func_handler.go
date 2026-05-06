package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"
	httpResponse "github.com/Benjamin-Gthub2/api-shared/custom-http/interfaces/rest"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

// GetAttendanceById is a method to get attendance by id
// @Summary Get attendance by id
// @Description Get attendance by id
// @Tags Attendances
// @Accept json
// @Produce json
// @Param attendanceId path string true "the id of the attendance"
// @Success 200 {object} attendanceByIdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/attendances/{attendanceId} [get]
// @Security BearerAuth
func (h attendancesHandler) GetAttendanceById(c *gin.Context) {
	ctx := c.Request.Context()
	attendanceId := c.Param("attendanceId")

	attendance, err := h.attendancesUseCase.GetAttendanceById(ctx, attendanceId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := attendanceByIdResult{
		Data:   attendance,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// GetAttendances is a method to get attendances
// @Summary Get attendances
// @Description Get attendances
// @Tags Attendances
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size_page query int false "Size page"
// @Param start_date query string false "Start date filter"
// @Param end_date query string false "End date filter"
// @Success 200 {object} attendancesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/attendances [get]
// @Security BearerAuth
func (h attendancesHandler) GetAttendances(c *gin.Context) {
	ctx := c.Request.Context()
	pagination := paramsDomain.NewPaginationParams(c.Request)
	searchParams := attendancesDomain.GetAttendancesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	attendances, paginationRes, err := h.attendancesUseCase.GetAttendances(ctx, pagination, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := attendancesResult{
		Data:       attendances,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateAttendance is a method to create an attendance
// @Summary Create attendance
// @Description Create attendance
// @Tags Attendances
// @Accept json
// @Produce json
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/attendances [post]
// @Security BearerAuth
func (h attendancesHandler) CreateAttendance(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetString("userId")

	id, err := h.attendancesUseCase.CreateAttendance(ctx, userId)
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

// DeleteAttendance is a method to delete an attendance
// @Summary Delete attendance
// @Description Delete attendance (soft delete)
// @Tags Attendances
// @Accept json
// @Produce json
// @Param attendanceId path string true "the id of the attendance"
// @Success 200 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/event/attendances/{attendanceId} [delete]
// @Security BearerAuth
func (h attendancesHandler) DeleteAttendance(c *gin.Context) {
	ctx := c.Request.Context()
	attendanceId := c.Param("attendanceId")
	userId := c.GetString("userId")

	err := h.attendancesUseCase.DeleteAttendance(ctx, attendanceId, userId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   attendanceId,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
