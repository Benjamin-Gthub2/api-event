package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	eventsSharedDomain "github.com/Benjamin-Gthub2/api-event/events-shared/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendancesMySQLRepo struct {
	clock                  smartClock.Clock
	timeout                time.Duration
	err                    *errDomain.SmartError
	eventsSharedRepository eventsSharedDomain.EventSharedRepository
}

func NewAttendancesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
	eventsSharedRepository eventsSharedDomain.EventSharedRepository,
) attendancesDomain.AttendancesRepository {
	rep := &attendancesMySQLRepo{
		clock:                  clock,
		timeout:                time.Duration(mongoTimeout) * time.Second,
		err:                    errDomain.NewErr().SetLayer(errDomain.Infra),
		eventsSharedRepository: eventsSharedRepository,
	}
	return rep
}
