package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendancesUseCase struct {
	attendancesRepository   attendancesDomain.AttendancesRepository
	attendancesRTRepository attendancesDomain.AttendancesRTRepository
	validationRepository    validationsDomain.ValidationRepository
	authRepository          authDomain.AuthRepository
	contextTimeout          time.Duration
	err                     *errDomain.SmartError
}

func NewAttendancesUseCase(
	ur attendancesDomain.AttendancesRepository,
	attendancesRTRepository attendancesDomain.AttendancesRTRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) attendancesDomain.AttendancesUseCase {
	return &attendancesUseCase{
		attendancesRepository:   ur,
		attendancesRTRepository: attendancesRTRepository,
		validationRepository:    validation,
		authRepository:          authRepository,
		contextTimeout:          timeout,
		err:                     errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
