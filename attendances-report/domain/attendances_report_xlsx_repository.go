package domain

import (
	"bytes"
	"context"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type AttendancesXlsxRepository interface {
	GenerateAttendancesReportXlsx(
		ctx context.Context,
		attendances []attendancesDomain.Attendance,
		displayFilters AttendancesReportDisplayFilters,
	) (*bytes.Buffer, error)
}
