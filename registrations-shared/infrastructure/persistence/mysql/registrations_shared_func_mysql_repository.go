package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	"github.com/stroiman/go-automapper"

	registrationSharedDomain "github.com/Benjamin-Gthub2/api-event/registrations-shared/domain"
)

//go:embed sql/get_registration_status_by_code.sql
var QueryGetRegistrationStatusByCode string

func (r registrationSharedMySQLRepo) GetStatusByCode(
	ctx context.Context,
	code string,
) (
	registrationStatusById *registrationSharedDomain.RegistrationStatus,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStatusByCode").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetRegistrationStatusByCode, code)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStatusByCode").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	registrationStatusesTmp := make([]RegistrationStatusRow, 0)
	err = carta.Map(results, &registrationStatusesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetStatusByCode").SetRaw(err)
	}

	registrationStatusesAux := make([]registrationSharedDomain.RegistrationStatus, 0)
	automapper.Map(registrationStatusesTmp, &registrationStatusesAux)
	if len(registrationStatusesAux) == 0 {
		return nil, registrationSharedDomain.ErrRegistrationStatusNotFound
	}
	return &registrationStatusesAux[0], nil
}
