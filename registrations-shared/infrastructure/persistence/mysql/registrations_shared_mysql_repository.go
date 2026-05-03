package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"

	registrationSharedDomain "github.com/Benjamin-Gthub2/api-event/registrations-shared/domain"
)

type registrationSharedMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRegistrationSharedRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) registrationSharedDomain.RegistrationSharedRepository {
	rep := &registrationSharedMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
