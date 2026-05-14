package xlsx

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
)

type attendancesReportXlsxRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewAttendancesReportXlsxRepository(
	clock smartClock.Clock,
	timeout int,
) domain.AttendancesXlsxRepository {
	return &attendancesReportXlsxRepo{
		clock:   clock,
		timeout: time.Duration(timeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
}
