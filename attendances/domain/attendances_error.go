package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrAttendanceNotFoundCode          = "ERR_ATTENDANCE_NOT_FOUND"
	ErrWorkshopsNotFoundCode           = "ERR_WORKSHOP_NOT_FOUND"
	ErrPeopleNotFoundCode              = "ERR_PERSON_NOT_FOUND"
	ErrAttendanceAlreadyExistsCode     = "ERR_ATTENDANCE_ALREADY_EXISTS"
	ErrAttendanceScheduleConflictCode  = "ERR_ATTENDANCE_SCHEDULE_CONFLICT"
)

var (
	ErrAttendanceNotFound = errDomain.NewErr().
				SetCode(ErrAttendanceNotFoundCode).
				SetDescription("ATTENDANCE NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("GetAttendanceById")
	ErrWorkshopNotFound = errDomain.NewErr().
				SetCode(ErrWorkshopsNotFoundCode).
				SetDescription("WORKSHOP NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateAttendance")
	ErrPersonNotFound = errDomain.NewErr().
				SetCode(ErrPeopleNotFoundCode).
				SetDescription("PERSON NOT FOUND").
				SetLevel(errDomain.LevelError).
				SetHttpStatus(http.StatusNotFound).
				SetLayer(errDomain.UseCase).
				SetFunction("CreateRegistration")
	ErrAttendanceAlreadyExists = errDomain.NewErr().
					SetCode(ErrAttendanceAlreadyExistsCode).
					SetDescription("THE BENEFICIARY IS ALREADY REGISTERED IN THIS WORKSHOP").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateAttendance")
	ErrAttendanceScheduleConflict = errDomain.NewErr().
					SetCode(ErrAttendanceScheduleConflictCode).
					SetDescription("THE BENEFICIARY IS ALREADY REGISTERED IN ANOTHER WORKSHOP AT THE SAME START TIME").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusConflict).
					SetLayer(errDomain.UseCase).
					SetFunction("CreateAttendance")
)
