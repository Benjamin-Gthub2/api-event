package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

type workshopSpeakersMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewWorkshopSpeakersRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) workshopSpeakersDomain.WorkshopSpeakersRepository {
	rep := &workshopSpeakersMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
