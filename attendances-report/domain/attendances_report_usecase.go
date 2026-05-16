package domain

import (
	"bytes"
	"context"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type AttendancesReportDisplayFilters struct {
	EventName       string
	WorkshopName    string
	BeneficiaryName string
}

type AttendancesReportUseCase interface {
	GenerateAttendancesReportXlsx(
		ctx context.Context,
		searchParams attendancesDomain.GetAttendancesParams,
		displayFilters AttendancesReportDisplayFilters,
	) (*bytes.Buffer, error)
}
