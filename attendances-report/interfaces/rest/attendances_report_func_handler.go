package rest

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	restCore "github.com/Benjamin-Gthub2/api-shared/api-core/interfaces/rest"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
)

// GenerateAttendancesReportXlsx genera un reporte xlsx de asistencias
// @Summary Generate attendances xlsx report
// @Description Generate xlsx report with attendances filtered by query params
// @Tags AttendancesReport
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param event_id query string false "Event ID"
// @Param workshop_id query string false "Workshop ID"
// @Param beneficiary_id query string false "Beneficiary ID"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Param event_name query string false "Event name (used in title and filename)"
// @Param workshop_name query string false "Workshop name (used in title and filename)"
// @Param beneficiary_name query string false "Beneficiary name (used in title and filename)"
// @Success 200 {file} octet-stream
// @Failure 500 {object} errorDomain.SmartError "Internal Server Error"
// @Router /api/v1/event/attendances/xlsx_report [get]
// @Security BearerAuth
func (h attendancesReportHandler) GenerateAttendancesReportXlsx(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := attendancesDomain.GetAttendancesParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)

	displayFilters := domain.AttendancesReportDisplayFilters{
		EventName:       c.Query("event_name"),
		WorkshopName:    c.Query("workshop_name"),
		BeneficiaryName: c.Query("beneficiary_name"),
	}

	fileBin, err := h.attendancesReportUseCase.GenerateAttendancesReportXlsx(ctx, searchParams, displayFilters)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	restCore.Binary(c, http.StatusOK,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		fileBin.Bytes(),
		buildFilename(displayFilters),
	)
}

func buildFilename(filters domain.AttendancesReportDisplayFilters) string {
	parts := []string{"Reporte de Asistencias"}
	if filters.EventName != "" {
		parts = append(parts, filters.EventName)
	}
	if filters.WorkshopName != "" {
		parts = append(parts, filters.WorkshopName)
	}
	if filters.BeneficiaryName != "" {
		parts = append(parts, filters.BeneficiaryName)
	}
	return strings.Join(parts, " - ") + ".xlsx"
}
