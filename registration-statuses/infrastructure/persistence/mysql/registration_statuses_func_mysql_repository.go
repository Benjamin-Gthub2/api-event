package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/stroiman/go-automapper"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_registration_status_by_id.sql
var QueryGetRegistrationStatusById string

//go:embed sql/get_registration_statuses.sql
var QueryGetRegistrationStatuses string

//go:embed sql/get_total_registration_statuses.sql
var QueryGetTotalRegistrationStatuses string

//go:embed sql/create_registration_status.sql
var QueryCreateRegistrationStatus string

//go:embed sql/update_registration_status.sql
var QueryUpdateRegistrationStatus string

//go:embed sql/delete_registration_status.sql
var QueryDeleteRegistrationStatus string

func (r registrationStatusesMySQLRepo) GetRegistrationStatusById(
	ctx context.Context,
	registrationStatusId string,
) (
	registrationStatusById *registrationStatusesDomain.RegistrationStatus,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatusById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetRegistrationStatusById, registrationStatusId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatusById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	registrationStatusesTmp := make([]RegistrationStatus, 0)
	err = carta.Map(results, &registrationStatusesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatusById").SetRaw(err)
	}

	registrationStatusesAux := make([]registrationStatusesDomain.RegistrationStatus, 0)
	automapper.Map(registrationStatusesTmp, &registrationStatusesAux)
	if len(registrationStatusesAux) == 0 {
		return nil, registrationStatusesDomain.ErrRegistrationStatusNotFound
	}
	return &registrationStatusesAux[0], nil
}

func (r registrationStatusesMySQLRepo) GetRegistrationStatuses(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams registrationStatusesDomain.GetRegistrationStatusesParams,
) (
	registrationStatusesRows []registrationStatusesDomain.RegistrationStatus,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatuses").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetRegistrationStatuses,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatuses").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	registrationStatusesTmp := make([]RegistrationStatus, 0)
	err = carta.Map(results, &registrationStatusesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationStatuses").SetRaw(err)
	}

	var registrationStatuses = make([]registrationStatusesDomain.RegistrationStatus, 0)
	automapper.Map(registrationStatusesTmp, &registrationStatuses)
	return registrationStatuses, nil
}

func (r registrationStatusesMySQLRepo) GetTotalRegistrationStatuses(
	ctx context.Context,
	searchParams registrationStatusesDomain.GetRegistrationStatusesParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrationStatuses").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalRegistrationStatuses,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrationStatuses").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r registrationStatusesMySQLRepo) CreateRegistrationStatus(
	ctx context.Context,
	body registrationStatusesDomain.CreateRegistrationStatus,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRegistrationStatus").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateRegistrationStatus,
		body.Id,
		body.Code,
		body.Description,
		body.Position,
		body.Enable,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRegistrationStatus").SetRaw(err)
	}
	return
}

func (r registrationStatusesMySQLRepo) UpdateRegistrationStatus(
	ctx context.Context,
	body registrationStatusesDomain.UpdateRegistrationStatus,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationStatus").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateRegistrationStatus,
		body.Code,
		body.Description,
		body.Position,
		body.Enable,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateRegistrationStatus").SetRaw(err)
	}
	return
}

func (r registrationStatusesMySQLRepo) DeleteRegistrationStatus(
	ctx context.Context,
	body registrationStatusesDomain.DeleteRegistrationStatus,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteRegistrationStatus").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteRegistrationStatus,
		now,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteRegistrationStatus").SetRaw(err)
	}
	return
}
