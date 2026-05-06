package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

type registrationStatusesMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRegistrationStatusesRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) registrationStatusesDomain.RegistrationStatusesRepository {
	rep := &registrationStatusesMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
