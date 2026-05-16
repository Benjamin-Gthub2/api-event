/*
 * File: events_func_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains functions for interacting with the repository layer related to events.
 *
 * Last Modified: 2026-04-15
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/Benjamin-Gthub2/carta"
	"github.com/stroiman/go-automapper"

	"github.com/Benjamin-Gthub2/api-shared/db"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	eventDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_total_events.sql
var QueryGetTotalEvents string

//go:embed sql/get_events.sql
var QueryGetEvents string

//go:embed sql/update_event.sql
var QueryUpdateEvent string

//go:embed sql/delete_event.sql
var QueryDeleteEvent string

//go:embed sql/create_event.sql
var QueryCreateEvent string

//go:embed sql/get_roles_users_by_event.sql
var QueryGetRolesUsersByEvent string

//go:embed sql/enable_disable_event.sql
var QueryEnableDisableEvent string

//go:embed sql/get_event_sums.sql
var QueryGetEventSums string

func (r eventsMySQLRepo) GetEvents(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	params eventDomain.GetEventsParams,
) (
	eventsRows []eventDomain.Event,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEvents").SetRaw(err)
	}
	results, err := client.QueryContext(ctx,
		QueryGetEvents,
		params.Status,
		params.Status,
		params.NameOrDocument,
		params.NameOrDocument,
		params.NameOrDocument,
		sizePage,
		offset)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEvents").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	eventTmp := make([]EventHelper, 0)
	err = carta.Map(results, &eventTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEvents").SetRaw(err)
	}
	automapper.Map(eventTmp, &eventsRows)

	//for iEvent, event := range eventsRows {
	//	if event.EventFiles[0].Id == nil {
	//		eventsRows[iEvent].EventFiles = []eventDomain.EventFile{}
	//	}
	//}

	return eventsRows, nil
}

func (r eventsMySQLRepo) GetTotalEvents(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	params eventDomain.GetEventsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEvents").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalEvents,
			params.Status,
			params.Status,
			params.NameOrDocument,
			params.NameOrDocument,
			params.NameOrDocument,
		).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEvents").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r eventsMySQLRepo) CreateEvent(
	ctx context.Context,
	eventId string,
	body eventDomain.CreateEventBody,
) (
	lastId *string,
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateEvent").SetRaw(err)
	}
	tx, err = client.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")

	_, err = tx.ExecContext(ctx,
		QueryCreateEvent,
		eventId,
		body.Name,
		body.Description,
		body.Code,
		body.Phone,
		body.Document,
		body.Address,
		body.Industry,
		body.Enable,
		now)

	if err != nil {
		return nil, r.err.Clone().SetFunction("CreateEvent").SetRaw(err)
	}

	//err = r.roleRepository.CreateRoleBatch(ctx, tx, defaultToEvent.NewRoles)
	//if err != nil {
	//	return nil, err
	//}
	//err = r.policyRepository.CreatePoliciesBatch(ctx, tx, defaultToEvent.NewPolicies)
	//if err != nil {
	//	return nil, err
	//}
	//err = r.rolePoliciesRepository.CreateRolePoliciesBatch(ctx, tx, defaultToEvent.NewRolePolicies)
	//if err != nil {
	//	return nil, err
	//}
	//err = r.policyPermissionsRepository.CreatePolicyPermissionsBatch(ctx, tx, defaultToEvent.NewPolicyPermissions)
	//if err != nil {
	//	return nil, err
	//}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	lastId = &eventId
	return
}

func (r eventsMySQLRepo) UpdateEvent(
	ctx context.Context,
	eventId string,
	body eventDomain.UpdateEventBody,
) (
	err error,
) {
	var tx *sql.Tx
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEvent").SetRaw(err)
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
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	_, err = tx.ExecContext(
		ctx,
		QueryUpdateEvent,
		body.Name,
		body.Description,
		body.Code,
		body.Phone,
		body.Document,
		body.Address,
		body.Industry,
		body.Enable,
		eventId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEvent").SetRaw(err)
	}

	//err = r.roleRepository.CreateRoleBatch(ctx, tx, defaultToEvent.NewRoles)
	//if err != nil {
	//	return err
	//}
	//err = r.policyRepository.CreatePoliciesBatch(ctx, tx, defaultToEvent.NewPolicies)
	//if err != nil {
	//	return err
	//}
	//err = r.rolePoliciesRepository.CreateRolePoliciesBatch(ctx, tx, defaultToEvent.NewRolePolicies)
	//if err != nil {
	//	return err
	//}
	//err = r.policyPermissionsRepository.CreatePolicyPermissionsBatch(ctx, tx, defaultToEvent.NewPolicyPermissions)
	//if err != nil {
	//	return err
	//}

	if err = tx.Commit(); err != nil {
		return err
	}
	return
}

func (r eventsMySQLRepo) DeleteEvent(
	ctx context.Context,
	id string,
) (
	updated bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteEvent").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryDeleteEvent,
		now,
		id)

	if err != nil {
		return false, r.err.Clone().SetFunction("DeleteEvent").SetRaw(err)
	}
	return true, nil
}

func (r eventsMySQLRepo) GetRolesByEvent(
	ctx context.Context,
	eventId string,
) (
	roles []eventDomain.Role,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRolesByEvent").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetRolesUsersByEvent, eventId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRolesByEvent").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	RolesTmp := make([]RoleByEvent, 0)
	err = carta.Map(results, &RolesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRolesByEvent").SetRaw(err)
	}
	for iRole, role := range RolesTmp {
		usersTmp := make([]User, 0)
		for _, user := range role.Users {
			if user.Id != nil {
				usersTmp = append(usersTmp, user)
			}
		}
		RolesTmp[iRole].Users = usersTmp
	}
	automapper.Map(RolesTmp, &roles)
	return roles, nil
}

func (r eventsMySQLRepo) EnableDisableEvent(
	ctx context.Context,
	eventId string,
	body eventDomain.EnableDisableEventRequest,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("EnableDisableEvent").SetRaw(err)
	}
	_, err = client.ExecContext(
		ctx,
		QueryEnableDisableEvent,
		body.Enable,
		eventId,
	)
	if err != nil {
		return r.err.Clone().SetFunction("EnableDisableEvent").SetRaw(err)
	}
	return
}

func (r eventsMySQLRepo) GetEventSums(
	ctx context.Context,
	params eventDomain.GetEventSumsParams,
) (
	eventsRows []eventDomain.EventSums,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventSums").SetRaw(err)
	}
	results, err := client.QueryContext(ctx,
		QueryGetEventSums,
		params.EventId,
		params.EventId,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventSums").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	eventTmp := make([]EventSums, 0)
	err = carta.Map(results, &eventTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventSums").SetRaw(err)
	}
	automapper.Map(eventTmp, &eventsRows)

	return eventsRows, nil
}
