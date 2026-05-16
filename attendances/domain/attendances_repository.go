package domain

import (
	"context"
	"database/sql"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type AttendancesRepository interface {
	GetAttendanceById(ctx context.Context, attendanceId string) (*Attendance, error)
	GetAttendances(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetAttendancesParams) ([]Attendance, error)
	GetTotalAttendances(ctx context.Context, searchParams GetAttendancesParams) (*int, error)
	AttendanceExistsByWorkshopAndBeneficiary(ctx context.Context, workshopId, beneficiaryId string) (bool, error)
	AttendanceExistsByBeneficiaryAndStartDate(ctx context.Context, beneficiaryId, workshopId string) (bool, error)
	CreateAttendance(ctx context.Context, tx *sql.Tx, body CreateAttendance) error
	MainCreateAttendance(ctx context.Context, body CreateAttendance) (err error)
	DeleteAttendance(ctx context.Context, tx *sql.Tx, body DeleteAttendance) error
	MainDeleteAttendance(ctx context.Context, body DeleteAttendance) (err error)
}
