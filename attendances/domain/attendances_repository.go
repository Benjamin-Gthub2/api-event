package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type AttendancesRepository interface {
	GetAttendanceById(ctx context.Context, attendanceId string) (*Attendance, error)
	GetAttendances(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetAttendancesParams) ([]Attendance, error)
	GetTotalAttendances(ctx context.Context, searchParams GetAttendancesParams) (*int, error)
	CreateAttendance(ctx context.Context, body CreateAttendance) error
	DeleteAttendance(ctx context.Context, body DeleteAttendance) error
}
