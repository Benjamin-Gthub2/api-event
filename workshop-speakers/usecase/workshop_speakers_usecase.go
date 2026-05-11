package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

type workshopSpeakersUseCase struct {
	workshopSpeakersRepository workshopSpeakersDomain.WorkshopSpeakersRepository
	validationRepository       validationsDomain.ValidationRepository
	authRepository             authDomain.AuthRepository
	contextTimeout             time.Duration
	err                        *errDomain.SmartError
}

func NewWorkshopSpeakersUseCase(
	ur workshopSpeakersDomain.WorkshopSpeakersRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) workshopSpeakersDomain.WorkshopSpeakersUseCase {
	return &workshopSpeakersUseCase{
		workshopSpeakersRepository: ur,
		validationRepository:       validation,
		authRepository:             authRepository,
		contextTimeout:             timeout,
		err:                        errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
