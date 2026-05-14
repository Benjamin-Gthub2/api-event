package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
)

type attendancesReportUseCase struct {
	attendancesXlsxRepository domain.AttendancesXlsxRepository
	attendancesRepository     attendancesDomain.AttendancesRepository
	authRepository            authDomain.AuthRepository
	contextTimeout            time.Duration
	err                       *errDomain.SmartError
}

func NewAttendancesReportUseCase(
	attendancesXlsxRepository domain.AttendancesXlsxRepository,
	attendancesRepository attendancesDomain.AttendancesRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.AttendancesReportUseCase {
	return &attendancesReportUseCase{
		attendancesXlsxRepository: attendancesXlsxRepository,
		attendancesRepository:     attendancesRepository,
		authRepository:            authRepository,
		contextTimeout:            timeout,
		err:                       errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
