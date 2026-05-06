package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

type eventTypesUseCase struct {
	eventTypesRepository eventTypesDomain.EventTypesRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewEventTypesUseCase(
	ur eventTypesDomain.EventTypesRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) eventTypesDomain.EventTypesUseCase {
	return &eventTypesUseCase{
		eventTypesRepository: ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
