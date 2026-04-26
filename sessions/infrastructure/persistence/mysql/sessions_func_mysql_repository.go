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

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

//go:embed sql/get_session_by_id.sql
var QueryGetSessionById string

//go:embed sql/get_sessions.sql
var QueryGetSessions string

//go:embed sql/get_total_sessions.sql
var QueryGetTotalSessions string

//go:embed sql/create_session.sql
var QueryCreateSession string

//go:embed sql/update_session.sql
var QueryUpdateSession string

//go:embed sql/delete_session.sql
var QueryDeleteSession string

//go:embed sql/get_session_sums.sql
var QueryGetSessionSums string

func (r sessionsMySQLRepo) GetSessionById(
	ctx context.Context,
	sessionId string,
) (
	sessionById *sessionsDomain.Session,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionById").SetRaw(err)
	}
	results, err := client.QueryContext(ctx, QueryGetSessionById, sessionId)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	sessionsTmp := make([]Session, 0)
	err = carta.Map(results, &sessionsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionById").SetRaw(err)
	}

	sessionsAux := make([]sessionsDomain.Session, 0)
	automapper.Map(sessionsTmp, &sessionsAux)
	if len(sessionsAux) == 0 {
		return nil, sessionsDomain.ErrSessionNotFound
	}
	return &sessionsAux[0], nil
}

func (r sessionsMySQLRepo) GetSessions(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams sessionsDomain.GetSessionsParams,
) (
	sessionsRows []sessionsDomain.Session,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessions").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetSessions,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.EndDate,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessions").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	sessionsTmp := make([]Session, 0)
	err = carta.Map(results, &sessionsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessions").SetRaw(err)
	}

	var sessions = make([]sessionsDomain.Session, 0)
	automapper.Map(sessionsTmp, &sessions)
	return sessions, nil
}

func (r sessionsMySQLRepo) GetTotalSessions(
	ctx context.Context,
	searchParams sessionsDomain.GetSessionsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalSessions").SetRaw(err)
	}
	err = client.QueryRowContext(
		ctx,
		QueryGetTotalSessions,
		searchParams.WorkshopId,
		searchParams.WorkshopId,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.EndDate,
	).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalSessions").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r sessionsMySQLRepo) CreateSession(
	ctx context.Context,
	body sessionsDomain.CreateSession,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateSession").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateSession,
		body.Id,
		body.WorkshopId,
		body.StartDate,
		body.EndDate,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateSession").SetRaw(err)
	}
	return
}

func (r sessionsMySQLRepo) UpdateSession(
	ctx context.Context,
	body sessionsDomain.UpdateSession,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateSession").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryUpdateSession,
		body.StartDate,
		body.EndDate,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("UpdateSession").SetRaw(err)
	}
	return
}

func (r sessionsMySQLRepo) DeleteSession(
	ctx context.Context,
	body sessionsDomain.DeleteSession,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:05")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteSession").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryDeleteSession,
		now,
		body.DeletedBy,
		body.Id,
	)
	if err != nil {
		return r.err.Clone().SetFunction("DeleteSession").SetRaw(err)
	}
	return
}

func (r sessionsMySQLRepo) GetSessionSums(
	ctx context.Context,
	searchParams sessionsDomain.GetSessionSumsParams,
) (
	sessionsRows []sessionsDomain.SessionSums,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionSums").SetRaw(err)
	}
	results, err := client.QueryContext(ctx,
		QueryGetSessionSums,
		searchParams.SessionId,
		searchParams.SessionId,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionSums").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)
	sessionTmp := make([]SessionSums, 0)
	err = carta.Map(results, &sessionTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetSessionSums").SetRaw(err)
	}
	automapper.Map(sessionTmp, &sessionsRows)

	return sessionsRows, nil
}
