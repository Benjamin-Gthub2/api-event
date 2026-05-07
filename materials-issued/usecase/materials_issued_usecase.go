package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

type materialsIssuedUseCase struct {
	materialsIssuedRepository materialsIssuedDomain.MaterialsIssuedRepository
	validationRepository      validationsDomain.ValidationRepository
	authRepository            authDomain.AuthRepository
	contextTimeout            time.Duration
	err                       *errDomain.SmartError
}

func NewMaterialsIssuedUseCase(
	ur materialsIssuedDomain.MaterialsIssuedRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) materialsIssuedDomain.MaterialsIssuedUseCase {
	return &materialsIssuedUseCase{
		materialsIssuedRepository: ur,
		validationRepository:      validation,
		authRepository:            authRepository,
		contextTimeout:            timeout,
		err:                       errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
