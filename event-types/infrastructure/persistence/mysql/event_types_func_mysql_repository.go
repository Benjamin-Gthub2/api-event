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

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

//go:embed sql/get_event_type_by_id.sql
var QueryGetEventTypeById string

//go:embed sql/get_event_types.sql
var QueryGetEventTypes string

//go:embed sql/get_total_event_types.sql
var QueryGetTotalEventTypes string

//go:embed sql/create_event_type.sql
var QueryCreateEventType string

//go:embed sql/update_event_type.sql
var QueryUpdateEventType string

//go:embed sql/delete_event_type.sql
var QueryDeleteEventType string

func (r eventTypesMySQLRepo) GetEventTypeById(
	ctx context.Context,
	eventTypeId string,
) (
	eventTypeById *eventTypesDomain.EventType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypeById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetEventTypeById, eventTypeId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypeById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	eventTypesTmp := make([]EventType, 0)
	err = carta.Map(results, &eventTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypeById").SetRaw(err)
	}

	eventTypesAux := make([]eventTypesDomain.EventType, 0)
	automapper.Map(eventTypesTmp, &eventTypesAux)
	if len(eventTypesAux) == 0 {
		return nil, eventTypesDomain.ErrEventTypeNotFound
	}
	return &eventTypesAux[0], nil
}

func (r eventTypesMySQLRepo) GetEventTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams eventTypesDomain.GetEventTypesParams,
) (
	eventTypesRows []eventTypesDomain.EventType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypes").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetEventTypes,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypes").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	eventTypesTmp := make([]EventType, 0)
	err = carta.Map(results, &eventTypesTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetEventTypes").SetRaw(err)
	}

	var eventTypes = make([]eventTypesDomain.EventType, 0)
	automapper.Map(eventTypesTmp, &eventTypes)
	return eventTypes, nil
}

func (r eventTypesMySQLRepo) GetTotalEventTypes(
	ctx context.Context,
	searchParams eventTypesDomain.GetEventTypesParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEventTypes").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalEventTypes,
		searchParams.SearchValue, searchParams.SearchValue, searchParams.SearchValue,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalEventTypes").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r eventTypesMySQLRepo) CreateEventType(
	ctx context.Context,
	body eventTypesDomain.CreateEventType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateEventType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateEventType,
		body.Id,
		body.Code,
		body.Description,
		body.Enable,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateEventType").SetRaw(err)
	}
	return
}

func (r eventTypesMySQLRepo) UpdateEventType(
	ctx context.Context,
	body eventTypesDomain.UpdateEventType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEventType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateEventType,
		body.Code,
		body.Description,
		body.Enable,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateEventType").SetRaw(err)
	}
	return
}

func (r eventTypesMySQLRepo) DeleteEventType(
	ctx context.Context,
	body eventTypesDomain.DeleteEventType,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().In(limaLoc).Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteEventType").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteEventType,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteEventType").SetRaw(err)
	}
	return
}
