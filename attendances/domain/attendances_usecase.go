package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type AttendancesUseCase interface {
	GetAttendanceById(ctx context.Context, attendanceId string) (*Attendance, error)
	GetAttendances(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetAttendancesParams) ([]Attendance, *paramsDomain.PaginationResults, error)
	CreateAttendance(ctx context.Context, userId string) (*string, error)
	DeleteAttendance(ctx context.Context, attendanceId string, userId string) error
}
