package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/smart0n3/api-shared/db"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	"github.com/stroiman/go-automapper"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

//go:embed sql/get_workshop_type_by_id.sql
var QueryGetWorkshopTypeById string

//go:embed sql/get_workshop_types.sql
var QueryGetWorkshopTypes string

//go:embed sql/get_total_workshop_types.sql
var QueryGetTotalWorkshopTypes string

//go:embed sql/create_workshop_type.sql
var QueryCreateWorkshopType string

//go:embed sql/update_workshop_type.sql
var QueryUpdateWorkshopType string

//go:embed sql/delete_workshop_type.sql
var QueryDeleteWorkshopType string

func (r workshopTypesMySQLRepo) GetWorkshopTypeById(
	ctx context.Context,
	workshopTypeId string,
) (
	workshopTypeById *workshopTypesDomain.WorkshopType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypeById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetWorkshopTypeById, workshopTypeId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypeById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopTypesTmp := make([]WorkshopType, 0)
	err = carta.Map(results, &workshopTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypeById").SetRaw(err)
	}

	workshopTypesAux := make([]workshopTypesDomain.WorkshopType, 0)
	automapper.Map(workshopTypesTmp, &workshopTypesAux)
	if len(workshopTypesAux) == 0 {
		return nil, workshopTypesDomain.ErrWorkshopTypeNotFound
	}
	return &workshopTypesAux[0], nil
}

func (r workshopTypesMySQLRepo) GetWorkshopTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopTypesDomain.GetWorkshopTypesParams,
) (
	workshopTypesRows []workshopTypesDomain.WorkshopType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypes").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetWorkshopTypes,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopTypesTmp := make([]WorkshopType, 0)
	err = carta.Map(results, &workshopTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetWorkshopTypes").SetRaw(err)
	}

	var workshopTypes = make([]workshopTypesDomain.WorkshopType, 0)
	automapper.Map(workshopTypesTmp, &workshopTypes)
	return workshopTypes, nil
}

func (r workshopTypesMySQLRepo) GetTotalWorkshopTypes(
	ctx context.Context,
	searchParams workshopTypesDomain.GetWorkshopTypesParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshopTypes").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalWorkshopTypes,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalWorkshopTypes").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r workshopTypesMySQLRepo) CreateWorkshopType(
	ctx context.Context,
	body workshopTypesDomain.CreateWorkshopType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshopType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateWorkshopType,
		body.Id,
		body.Code,
		body.Description,
		body.Enable,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateWorkshopType").SetRaw(err)
	}
	return
}

func (r workshopTypesMySQLRepo) UpdateWorkshopType(
	ctx context.Context,
	body workshopTypesDomain.UpdateWorkshopType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateWorkshopType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateWorkshopType,
		body.Code,
		body.Description,
		body.Enable,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateWorkshopType").SetRaw(err)
	}
	return
}

func (r workshopTypesMySQLRepo) DeleteWorkshopType(
	ctx context.Context,
	body workshopTypesDomain.DeleteWorkshopType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshopType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteWorkshopType,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteWorkshopType").SetRaw(err)
	}
	return
}
