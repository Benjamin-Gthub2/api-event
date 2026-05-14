package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	eventSharedDomain "github.com/Benjamin-Gthub2/api-event/events-shared/domain"
	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

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

func intToPtr(value int) *int {
	return &value
}

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
		searchParams.EventId,
		searchParams.EventId,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.BeneficiaryId,
		searchParams.BeneficiaryId,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
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
		searchParams.EventId,
		searchParams.EventId,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.BeneficiaryId,
		searchParams.BeneficiaryId,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalAttendances").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r attendancesMySQLRepo) MainCreateAttendance(
	ctx context.Context,
	body attendancesDomain.CreateAttendance,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var tx *sql.Tx

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("MainCreateAttendance").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var sessionWorkshopEventById *eventSharedDomain.EventWorkshopSession
	sessionWorkshopEventById, err = r.eventsSharedRepository.GetSessionWorkshopEventByWorkshopId(ctx, tx, body.WorkshopId)
	if err != nil {
		return err
	}

	eventId := sessionWorkshopEventById.EventId
	workshopId := sessionWorkshopEventById.WorkshopId

	var eventTotals *eventSharedDomain.EventTotals
	var workshopTotals *eventSharedDomain.WorkshopTotals
	var updateEventTotals eventSharedDomain.UpdateEventTotals
	var updateWorkshopTotals eventSharedDomain.UpdateWorkshopTotals

	eventTotals, err = r.eventsSharedRepository.GetEventTotals(ctx, tx, eventId)
	if err != nil {
		return err
	}
	updateEventTotals = eventSharedDomain.UpdateEventTotals{
		TotalPres: intToPtr(eventTotals.TotalPres + 1),
	}
	err = r.eventsSharedRepository.UpdateEventTotals(ctx, tx, eventId, updateEventTotals)
	if err != nil {
		return err
	}

	workshopTotals, err = r.eventsSharedRepository.GetWorkshopTotals(ctx, tx, workshopId)
	if err != nil {
		return err
	}
	updateWorkshopTotals = eventSharedDomain.UpdateWorkshopTotals{
		TotalPres: intToPtr(workshopTotals.TotalPres + 1),
	}
	err = r.eventsSharedRepository.UpdateWorkshopTotals(ctx, tx, workshopId, updateWorkshopTotals)
	if err != nil {
		return err
	}

	err = r.CreateAttendance(ctx, tx, body)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return err
}

func (r attendancesMySQLRepo) CreateAttendance(
	ctx context.Context,
	tx *sql.Tx,
	body attendancesDomain.CreateAttendance,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	_, err = tx.ExecContext(ctx,
		QueryCreateAttendance,
		body.Id,
		body.WorkshopId,
		body.BeneficiaryId,
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
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
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
