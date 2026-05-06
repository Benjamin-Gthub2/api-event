package jwt

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	"github.com/stroiman/go-automapper"

	"github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

//go:embed sql/get_merchant_stores.sql
var QueryGetMerchantsStores string

//go:embed sql/get_merchant_stores_postgres.sql
var QueryGetMerchantsStoresPostgres string

func (a authJWTRepo) GetMerchantStoresByUser(
	ctx context.Context,
	userId string,
) (
	modules []domain.ModuleMid,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("GetMerchantStoresByUser")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	query := QueryGetMerchantsStores
	if dbDriver == "postgres" {
		query = QueryGetMerchantsStoresPostgres
	}
	results, err := client.QueryContext(
		ctx,
		query,
		userId,
	)
	if err != nil {
		return nil, a.err.Clone().SetFunction("GetMerchantStoresByUser").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	modulesTmp := make([]ModuleMid, 0)
	err = carta.Map(results, &modulesTmp)
	if err != nil {
		return nil, a.err.Clone().SetFunction("GetMerchantStoresByUser").SetRaw(err)
	}
	modules = make([]domain.ModuleMid, 0)
	automapper.Map(modulesTmp, &modules)
	return modules, nil
}
