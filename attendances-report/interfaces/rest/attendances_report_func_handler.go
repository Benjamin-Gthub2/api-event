package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

// GenerateAttendancesReportXlsx generates an xlsx report of attendances
// @Summary Generate attendances xlsx report
// @Description Generate xlsx report with attendances filtered by query params
// @Tags AttendancesReport
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param event_id query string false "Event ID"
// @Param workshop_id query string false "Workshop ID"
// @Param beneficiary_id query string false "Beneficiary ID"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {file} octet-stream
// @Failure 500 {object} errorDomain.SmartError "Internal Server Error"
// @Router /api/v1/event/attendances/xlsx_report [get]
// @Security BearerAuth
func (h attendancesReportHandler) GenerateAttendancesReportXlsx(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := attendancesDomain.GetAttendancesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	fileBin, err := h.attendancesReportUseCase.GenerateAttendancesReportXlsx(ctx, searchParams)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	restCore.Binary(c, http.StatusOK,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		fileBin.Bytes(),
		"Reporte de Asistencias.xlsx",
	)
}
