package usecase

import (
	"time"

	authDomain "github.com/smart0n3/api-shared/auth/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

type registrationStatusesUseCase struct {
	registrationStatusesRepository registrationStatusesDomain.RegistrationStatusesRepository
	validationRepository           validationsDomain.ValidationRepository
	authRepository                 authDomain.AuthRepository
	contextTimeout                 time.Duration
	err                            *errDomain.SmartError
}

func NewRegistrationStatusesUseCase(
	ur registrationStatusesDomain.RegistrationStatusesRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) registrationStatusesDomain.RegistrationStatusesUseCase {
	return &registrationStatusesUseCase{
		registrationStatusesRepository: ur,
		validationRepository:           validation,
		authRepository:                 authRepository,
		contextTimeout:                 timeout,
		err:                            errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
