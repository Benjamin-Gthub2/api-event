package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

type sessionsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewSessionsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) sessionsDomain.SessionsRepository {
	rep := &sessionsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
