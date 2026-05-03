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

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

//go:embed sql/get_material_issued_by_id.sql
var QueryGetMaterialIssuedById string

//go:embed sql/get_materials_issued.sql
var QueryGetMaterialsIssued string

//go:embed sql/get_total_materials_issued.sql
var QueryGetTotalMaterialsIssued string

//go:embed sql/create_material_issued.sql
var QueryCreateMaterialIssued string

//go:embed sql/update_material_issued.sql
var QueryUpdateMaterialIssued string

//go:embed sql/delete_material_issued.sql
var QueryDeleteMaterialIssued string

func (r materialsIssuedMySQLRepo) GetMaterialIssuedById(
	ctx context.Context,
	materialIssuedId string,
) (
	materialIssuedById *materialsIssuedDomain.MaterialIssued,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialIssuedById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetMaterialIssuedById, materialIssuedId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialIssuedById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	materialsTmp := make([]MaterialIssued, 0)
	err = carta.Map(results, &materialsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialIssuedById").SetRaw(err)
	}

	materialsAux := make([]materialsIssuedDomain.MaterialIssued, 0)
	automapper.Map(materialsTmp, &materialsAux)
	if len(materialsAux) == 0 {
		return nil, materialsIssuedDomain.ErrMaterialIssuedNotFound
	}
	return &materialsAux[0], nil
}

func (r materialsIssuedMySQLRepo) GetMaterialsIssued(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams materialsIssuedDomain.GetMaterialsIssuedParams,
) (
	materialsRows []materialsIssuedDomain.MaterialIssued,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialsIssued").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetMaterialsIssued,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialsIssued").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	materialsTmp := make([]MaterialIssued, 0)
	err = carta.Map(results, &materialsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetMaterialsIssued").SetRaw(err)
	}

	var materials = make([]materialsIssuedDomain.MaterialIssued, 0)
	automapper.Map(materialsTmp, &materials)
	return materials, nil
}

func (r materialsIssuedMySQLRepo) GetTotalMaterialsIssued(
	ctx context.Context,
	searchParams materialsIssuedDomain.GetMaterialsIssuedParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalMaterialsIssued").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalMaterialsIssued,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalMaterialsIssued").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r materialsIssuedMySQLRepo) CreateMaterialIssued(
	ctx context.Context,
	body materialsIssuedDomain.CreateMaterialIssued,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateMaterialIssued").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateMaterialIssued,
		body.Id,
		body.Description,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateMaterialIssued").SetRaw(err)
	}
	return
}

func (r materialsIssuedMySQLRepo) UpdateMaterialIssued(
	ctx context.Context,
	body materialsIssuedDomain.UpdateMaterialIssued,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateMaterialIssued").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateMaterialIssued,
		body.Description,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateMaterialIssued").SetRaw(err)
	}
	return
}

func (r materialsIssuedMySQLRepo) DeleteMaterialIssued(
	ctx context.Context,
	body materialsIssuedDomain.DeleteMaterialIssued,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteMaterialIssued").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteMaterialIssued,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteMaterialIssued").SetRaw(err)
	}
	return
}
