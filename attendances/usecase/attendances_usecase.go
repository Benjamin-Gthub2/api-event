package usecase

import (
	"time"

	authDomain "github.com/smart0n3/api-shared/auth/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendancesUseCase struct {
	attendancesRepository attendancesDomain.AttendancesRepository
	validationRepository  validationsDomain.ValidationRepository
	authRepository        authDomain.AuthRepository
	contextTimeout        time.Duration
	err                   *errDomain.SmartError
}

func NewAttendancesUseCase(
	ur attendancesDomain.AttendancesRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) attendancesDomain.AttendancesUseCase {
	return &attendancesUseCase{
		attendancesRepository: ur,
		validationRepository:  validation,
		authRepository:        authRepository,
		contextTimeout:        timeout,
		err:                   errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
