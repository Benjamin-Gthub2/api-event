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

	"github.com/jackskj/carta"
	"github.com/stroiman/go-automapper"

	"github.com/smart0n3/api-shared/db"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"

	eventSharedDomain "github.com/Benjamin-Gthub2/api-event/events-shared/domain"
)

//go:embed sql/update_event_totals.sql
var QueryUpdateEventTotals string

//go:embed sql/update_workshop_totals.sql
var QueryUpdateWorkshopTotals string

//go:embed sql/update_session_totals.sql
var QueryUpdateSessionTotals string

//go:embed sql/get_event_totals.sql
var QueryGetEventTotals string

//go:embed sql/get_workshop_totals.sql
var QueryGetWorkshopTotals string

//go:embed sql/get_session_totals.sql
var QueryGetSessionTotals string

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

func (r eventSharedMySQLRepo) GetEventTotals(
	ctx context.Context,
	tx *sql.Tx,
	eventId string,
) (
	eventTotals *eventSharedDomain.EventTotals,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var results *sql.Rows
	if tx != nil {
		results, err = tx.QueryContext(
			ctx,
			QueryGetEventTotals,
			eventId,
		)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("GetEventTotals").SetRaw(err)
		}
		results, err = client.QueryContext(
			ctx,
			QueryGetEventTotals,
			eventId,
		)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	eventTotalsTmp := make([]EventTotals, 0)
	err = carta.Map(results, &eventTotalsTmp)
	if err != nil {
		return eventTotals, r.err.Clone().SetFunction("GetEventTotals").SetRaw(err)
	}

	eventTotalsAux := make([]eventSharedDomain.EventTotals, 0)
	automapper.Map(eventTotalsTmp, &eventTotalsAux)
	if len(eventTotalsAux) == 0 {
		return eventTotals, eventSharedDomain.ErrEventNotFound
	}
	return &eventTotalsAux[0], nil
}

func (r eventSharedMySQLRepo) GetWorkshopTotals(
	ctx context.Context,
	tx *sql.Tx,
	workshopId string,
) (
	workshopTotals *eventSharedDomain.WorkshopTotals,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var results *sql.Rows
	if tx != nil {
		results, err = tx.QueryContext(
			ctx,
			QueryGetEventTotals,
			workshopId,
		)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("GetWorkshopTotals").SetRaw(err)
		}
		results, err = client.QueryContext(
			ctx,
			QueryGetEventTotals,
			workshopId,
		)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	workshopTotalsTmp := make([]EventTotals, 0)
	err = carta.Map(results, &workshopTotalsTmp)
	if err != nil {
		return workshopTotals, r.err.Clone().SetFunction("GetWorkshopTotals").SetRaw(err)
	}

	workshopTotalsAux := make([]eventSharedDomain.WorkshopTotals, 0)
	automapper.Map(workshopTotalsTmp, &workshopTotalsAux)
	if len(workshopTotalsAux) == 0 {
		return workshopTotals, eventSharedDomain.ErrEventNotFound
	}
	return &workshopTotalsAux[0], nil
}

func (r eventSharedMySQLRepo) GetSessionTotals(
	ctx context.Context,
	tx *sql.Tx,
	sessionId string,
) (
	sessionTotals *eventSharedDomain.SessionTotals,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var results *sql.Rows
	if tx != nil {
		results, err = tx.QueryContext(
			ctx,
			QueryGetEventTotals,
			sessionId,
		)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("GetSessionTotals").SetRaw(err)
		}
		results, err = client.QueryContext(
			ctx,
			QueryGetEventTotals,
			sessionId,
		)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	sessionTotalsTmp := make([]EventTotals, 0)
	err = carta.Map(results, &sessionTotalsTmp)
	if err != nil {
		return sessionTotals, r.err.Clone().SetFunction("GetSessionTotals").SetRaw(err)
	}

	sessionTotalsAux := make([]eventSharedDomain.SessionTotals, 0)
	automapper.Map(sessionTotalsTmp, &sessionTotalsAux)
	if len(sessionTotalsAux) == 0 {
		return sessionTotals, eventSharedDomain.ErrEventNotFound
	}
	return &sessionTotalsAux[0], nil
}

func (r eventSharedMySQLRepo) GetSessionWorkshopEventById(
	ctx context.Context,
	tx *sql.Tx,
	registrationId string,
) (
	eventWorkshopSession *eventSharedDomain.EventWorkshopSession,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var results *sql.Rows
	if tx != nil {
		results, err = tx.QueryContext(
			ctx,
			QueryGetEventTotals,
			registrationId,
		)
	} else {
		var client *sql.DB
		client, _, err = db.ClientDB(ctx)
		if err != nil {
			return nil, r.err.Clone().SetFunction("GetSessionWorkshopEventById").SetRaw(err)
		}
		results, err = client.QueryContext(
			ctx,
			QueryGetEventTotals,
			registrationId,
		)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	eventWorkshopSessionTmp := make([]EventWorkshopSession, 0)
	err = carta.Map(results, &eventWorkshopSessionTmp)
	if err != nil {
		return eventWorkshopSession, r.err.Clone().SetFunction("GetSessionWorkshopEventById").SetRaw(err)
	}

	eventWorkshopSessionAux := make([]eventSharedDomain.EventWorkshopSession, 0)
	automapper.Map(eventWorkshopSessionTmp, &eventWorkshopSessionAux)
	if len(eventWorkshopSessionAux) == 0 {
		return eventWorkshopSession, eventSharedDomain.ErrEventNotFound
	}
	return &eventWorkshopSessionAux[0], nil
}
