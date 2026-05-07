package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

type eventTypesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewEventTypesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) eventTypesDomain.EventTypesRepository {
	rep := &eventTypesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
