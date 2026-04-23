package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

func (u sessionsUseCase) GetSessionById(
	ctx context.Context,
	sessionId string,
) (
	session *sessionsDomain.Session,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "sessions",
		IdColumnName:     "id",
		IdValue:          sessionId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sessionsDomain.ErrSessionNotFound
	}

	session, err = u.sessionsRepository.GetSessionById(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (u sessionsUseCase) GetSessions(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams sessionsDomain.GetSessionsParams,
) (
	res []sessionsDomain.Session,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetSessions, errGetTotalSessions error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetSessions, &wg)
		res, errGetSessions = u.sessionsRepository.GetSessions(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalSessions, &wg)
		total, errGetTotalSessions = u.sessionsRepository.GetTotalSessions(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetSessions != nil {
		return nil, nil, errGetSessions
	}
	if errGetTotalSessions != nil {
		return nil, nil, errGetTotalSessions
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u sessionsUseCase) CreateSession(
	ctx context.Context,
	userId string,
	body sessionsDomain.CreateSessionBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	sessionId := uuid.New().String()

	createSession := sessionsDomain.CreateSession{
		Id:         sessionId,
		WorkshopId: body.WorkshopId,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
		CreatedBy:  userId,
	}
	err = u.sessionsRepository.CreateSession(ctx, createSession)
	if err != nil {
		return nil, err
	}
	return &sessionId, nil
}

func (u sessionsUseCase) UpdateSession(
	ctx context.Context,
	sessionId string,
	body sessionsDomain.UpdateSessionBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "sessions",
		IdColumnName:     "id",
		IdValue:          sessionId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return sessionsDomain.ErrSessionNotFound
	}

	updateSession := sessionsDomain.UpdateSession{
		Id:        sessionId,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}
	err = u.sessionsRepository.UpdateSession(ctx, updateSession)
	if err != nil {
		return err
	}
	return nil
}

func (u sessionsUseCase) DeleteSession(
	ctx context.Context,
	sessionId string,
	userId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "sessions",
		IdColumnName:     "id",
		IdValue:          sessionId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return sessionsDomain.ErrSessionNotFound
	}

	deleteSession := sessionsDomain.DeleteSession{
		Id:        sessionId,
		DeletedBy: userId,
	}
	err = u.sessionsRepository.DeleteSession(ctx, deleteSession)
	if err != nil {
		return err
	}
	return nil
}
