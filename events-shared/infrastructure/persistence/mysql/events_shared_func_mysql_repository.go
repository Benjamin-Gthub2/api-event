/*
 * File: events_shared_func_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the function repository.
 *
 * Last Modified: 2026-04-24
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"

	eventSharedDomain "github.com/Benjamin-Gthub2/api-event/events-shared/domain"
)

//go:embed sql/update_event_totals.sql
var QueryUpdateEventTotals string

//go:embed sql/update_workshop_totals.sql
var QueryUpdateWorkshopTotals string

//go:embed sql/update_session_totals.sql
var QueryUpdateSessionTotals string

func (r eventSharedMySQLRepo) UpdateEventTotals(
	ctx context.Context,
	tx *sql.Tx,
	eventId string,
	body eventSharedDomain.UpdateEventTotals,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = tx.ExecContext(
		ctx,
		QueryUpdateEventTotals,
		body.TotalReg,
		body.TotalPay,
		body.TotalReg,
		eventId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEventTotals").SetRaw(err)
	}
	return
}

func (r eventSharedMySQLRepo) UpdateWorkshopTotals(
	ctx context.Context,
	tx *sql.Tx,
	eventId string,
	body eventSharedDomain.UpdateWorkshopTotals,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = tx.ExecContext(
		ctx,
		QueryUpdateWorkshopTotals,
		body.TotalReg,
		body.TotalPay,
		body.TotalReg,
		eventId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateWorkshopTotals").SetRaw(err)
	}
	return
}

func (r eventSharedMySQLRepo) UpdateSessionTotals(
	ctx context.Context,
	tx *sql.Tx,
	eventId string,
	body eventSharedDomain.UpdateSessionTotals,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	_, err = tx.ExecContext(
		ctx,
		QueryUpdateSessionTotals,
		body.TotalReg,
		body.TotalPay,
		body.TotalReg,
		eventId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateSessionTotals").SetRaw(err)
	}
	return
}
