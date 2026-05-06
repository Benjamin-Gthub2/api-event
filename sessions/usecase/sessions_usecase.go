package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

type sessionsUseCase struct {
	sessionsRepository   sessionsDomain.SessionsRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewSessionsUseCase(
	ur sessionsDomain.SessionsRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) sessionsDomain.SessionsUseCase {
	return &sessionsUseCase{
		sessionsRepository:   ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
