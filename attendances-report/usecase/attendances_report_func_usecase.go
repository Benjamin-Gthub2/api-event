package usecase

import (
	"bytes"
	"context"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
)

func (u attendancesReportUseCase) GenerateAttendancesReportXlsx(
	ctx context.Context,
	searchParams attendancesDomain.GetAttendancesParams,
	displayFilters domain.AttendancesReportDisplayFilters,
) (file *bytes.Buffer, err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	pagination := paramsDomain.PaginationParams{
		Page:     1,
		SizePage: 10000,
	}

	attendances, err := u.attendancesRepository.GetAttendances(ctx, pagination, searchParams)
	if err != nil {
		return nil, err
	}

	file, err = u.attendancesXlsxRepository.GenerateAttendancesReportXlsx(ctx, attendances, displayFilters)
	if err != nil {
		return nil, err
	}
	return file, nil
}
