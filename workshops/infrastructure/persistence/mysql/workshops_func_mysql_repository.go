package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_workshop_by_id.sql
var QueryGetWorkshopById string

//go:embed sql/get_workshops.sql
var QueryGetWorkshops string

//go:embed sql/get_total_workshops.sql
var QueryGetTotalWorkshops string

//go:embed sql/create_workshop.sql
var QueryCreateWorkshop string

//go:embed sql/update_workshop.sql
var QueryUpdateWorkshop string

//go:embed sql/delete_workshop.sql
var QueryDeleteWorkshop string

//go:embed sql/get_workshops_sums.sql
var QueryGetWorkshopsSums string

func (r workshopsMySQLRepo) GetWorkshopById(
	ctx context.Context,
	workshopId string,
) (
	workshopById *workshopsDomain.Workshop,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetWorkshopById, workshopId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopsTmp := make([]Workshop, 0)
	err = carta.Map(results, &workshopsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopById").SetRaw(err)
	}

	workshopsAux := make([]workshopsDomain.Workshop, 0)
	automapper.Map(workshopsTmp, &workshopsAux)
	if len(workshopsAux) == 0 {
		return nil, workshopsDomain.ErrWorkshopNotFound
	}
	return &workshopsAux[0], nil
}

func (r workshopsMySQLRepo) GetWorkshops(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopsDomain.GetWorkshopsParams,
) (
	workshopsRows []workshopsDomain.Workshop,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshops").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetWorkshops,
		searchParams.EventId,
		searchParams.EventId,
		searchParams.TypeId,
		searchParams.TypeId,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshops").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopsTmp := make([]Workshop, 0)
	err = carta.Map(results, &workshopsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshops").SetRaw(err)
	}

	var workshops = make([]workshopsDomain.Workshop, 0)
	automapper.Map(workshopsTmp, &workshops)
	return workshops, nil
}

func (r workshopsMySQLRepo) GetTotalWorkshops(
	ctx context.Context,
	searchParams workshopsDomain.GetWorkshopsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshops").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalWorkshops,
		searchParams.EventId,
		searchParams.EventId,
		searchParams.TypeId,
		searchParams.TypeId,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshops").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r workshopsMySQLRepo) CreateWorkshop(
	ctx context.Context,
	body workshopsDomain.CreateWorkshop,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshop").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateWorkshop,
		body.Id,
		body.TypeId,
		body.Name,
		body.Shortname,
		body.Code,
		body.Capacity,
		body.StartDate,
		body.EndDate,
		body.Place,
		body.EventId,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshop").SetRaw(err)
	}
	return
}

func (r workshopsMySQLRepo) UpdateWorkshop(
	ctx context.Context,
	body workshopsDomain.UpdateWorkshop,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateWorkshop").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateWorkshop,
		body.TypeId,
		body.Name,
		body.Shortname,
		body.Code,
		body.Capacity,
		body.StartDate,
		body.EndDate,
		body.Place,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateWorkshop").SetRaw(err)
	}
	return
}

func (r workshopsMySQLRepo) DeleteWorkshop(
	ctx context.Context,
	body workshopsDomain.DeleteWorkshop,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshop").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteWorkshop,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshop").SetRaw(err)
	}
	return
}

func (r workshopsMySQLRepo) GetWorkshopSums(
	ctx context.Context,
	searchParams workshopsDomain.GetWorkshopSumsParams,
) (
	workshopsRows []workshopsDomain.WorkshopSums,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSums").SetRaw(err)
	}
	results, err := client.QueryContext(ctx,
		QueryGetWorkshopsSums,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.SearchValue,
		searchParams.SearchValue,
		searchParams.SearchValue,
		searchParams.SearchValue,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.EndDate,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSums").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	workshopTmp := make([]WorkshopSums, 0)
	err = carta.Map(results, &workshopTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopSums").SetRaw(err)
	}
	automapper.Map(workshopTmp, &workshopsRows)

	return workshopsRows, nil
}
