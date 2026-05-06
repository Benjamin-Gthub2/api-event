package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

type workshopsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewWorkshopsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) workshopsDomain.WorkshopsRepository {
	rep := &workshopsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
