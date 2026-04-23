package usecase

import (
	"time"

	authDomain "github.com/smart0n3/api-shared/auth/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

type workshopTypesUseCase struct {
	workshopTypesRepository workshopTypesDomain.WorkshopTypesRepository
	validationRepository    validationsDomain.ValidationRepository
	authRepository          authDomain.AuthRepository
	contextTimeout          time.Duration
	err                     *errDomain.SmartError
}

func NewWorkshopTypesUseCase(
	ur workshopTypesDomain.WorkshopTypesRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) workshopTypesDomain.WorkshopTypesUseCase {
	return &workshopTypesUseCase{
		workshopTypesRepository: ur,
		validationRepository:    validation,
		authRepository:          authRepository,
		contextTimeout:          timeout,
		err:                     errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
