package domain

import (
	"bytes"
	"context"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type AttendancesReportUseCase interface {
	GenerateAttendancesReportXlsx(ctx context.Context, searchParams attendancesDomain.GetAttendancesParams) (*bytes.Buffer, error)
}
