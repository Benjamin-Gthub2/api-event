package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

type workshopTypesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewWorkshopTypesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) workshopTypesDomain.WorkshopTypesRepository {
	rep := &workshopTypesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
