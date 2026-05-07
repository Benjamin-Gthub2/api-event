package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

type workshopsUseCase struct {
	workshopsRepository  workshopsDomain.WorkshopsRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewWorkshopsUseCase(
	ur workshopsDomain.WorkshopsRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) workshopsDomain.WorkshopsUseCase {
	return &workshopsUseCase{
		workshopsRepository:  ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
