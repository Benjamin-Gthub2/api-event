package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/stroiman/go-automapper"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

//go:embed sql/get_attendance_by_id.sql
var QueryGetAttendanceById string

//go:embed sql/get_attendances.sql
var QueryGetAttendances string

//go:embed sql/get_total_attendances.sql
var QueryGetTotalAttendances string

//go:embed sql/create_attendance.sql
var QueryCreateAttendance string

//go:embed sql/delete_attendance.sql
var QueryDeleteAttendance string

func (r attendancesMySQLRepo) GetAttendanceById(
	ctx context.Context,
	attendanceId string,
) (
	attendanceById *attendancesDomain.Attendance,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendanceById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetAttendanceById, attendanceId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendanceById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	attendancesTmp := make([]Attendance, 0)
	err = carta.Map(results, &attendancesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendanceById").SetRaw(err)
	}

	attendancesAux := make([]attendancesDomain.Attendance, 0)
	automapper.Map(attendancesTmp, &attendancesAux)
	if len(attendancesAux) == 0 {
		return nil, attendancesDomain.ErrAttendanceNotFound
	}
	return &attendancesAux[0], nil
}

func (r attendancesMySQLRepo) GetAttendances(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams attendancesDomain.GetAttendancesParams,
) (
	attendancesRows []attendancesDomain.Attendance,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendances").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetAttendances,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendances").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	attendancesTmp := make([]Attendance, 0)
	err = carta.Map(results, &attendancesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetAttendances").SetRaw(err)
	}

	var attendances = make([]attendancesDomain.Attendance, 0)
	automapper.Map(attendancesTmp, &attendances)
	return attendances, nil
}

func (r attendancesMySQLRepo) GetTotalAttendances(
	ctx context.Context,
	searchParams attendancesDomain.GetAttendancesParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalAttendances").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalAttendances,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalAttendances").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r attendancesMySQLRepo) CreateAttendance(
	ctx context.Context,
	body attendancesDomain.CreateAttendance,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateAttendance").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateAttendance,
		body.Id,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateAttendance").SetRaw(err)
	}
	return
}

func (r attendancesMySQLRepo) DeleteAttendance(
	ctx context.Context,
	body attendancesDomain.DeleteAttendance,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteAttendance").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteAttendance,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteAttendance").SetRaw(err)
	}
	return
}
