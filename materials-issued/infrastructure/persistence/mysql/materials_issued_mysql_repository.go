package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

type materialsIssuedMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewMaterialsIssuedRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) materialsIssuedDomain.MaterialsIssuedRepository {
	rep := &materialsIssuedMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
