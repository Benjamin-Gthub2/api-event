/*
 * File: events_func_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the event functions are located.
 *
 * Last Modified: 2026-04-15
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	eventsDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

func (u eventsUseCase) GetEvents(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	params eventsDomain.GetEventsParams,
) (
	res []eventsDomain.Event,
	paginationResults *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetEvents, errGetTotalEvents error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetEvents, &wg)
		res, errGetEvents = u.eventsRepository.GetEvents(ctx, pagination, params)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalEvents, &wg)
		total, errGetTotalEvents = u.eventsRepository.GetTotalEvents(ctx, pagination, params)
		wg.Done()
	}()
	wg.Wait()

	if errGetEvents != nil {
		return nil, nil, errGetEvents
	}
	if errGetTotalEvents != nil {
		return nil, nil, errGetTotalEvents
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u eventsUseCase) CreateEvent(
	ctx context.Context,
	body eventsDomain.CreateEventBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	deleted := "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_events",
		IdColumnName:     "document",
		IdValue:          body.Document,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, eventsDomain.ErrEventDocumentAlreadyExist
	}

	recordExistsCode := validationsDomain.RecordExistsParams{
		Table:            "core_events",
		IdColumnName:     "code",
		IdValue:          body.Code,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	existCode, err := u.validationRepository.RecordExists(ctx, recordExistsCode)
	if err != nil {
		return nil, err
	}
	if existCode {
		return nil, eventsDomain.ErrEventCodeAlreadyExist
	}

	eventId := uuid.New().String()

	// get roles by default
	//defaultToEvent, err := u.roleDefaultsRepository.GetRolesDefaults(ctx, &eventId, nil)
	//if err != nil {
	//	return nil, err
	//}
	id, err = u.eventsRepository.CreateEvent(ctx, eventId, body)
	return
}

func (u eventsUseCase) UpdateEvent(
	ctx context.Context,
	eventId string,
	body eventsDomain.UpdateEventBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	deleted := "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_events",
		IdColumnName:     "id",
		IdValue:          eventId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(eventsDomain.ErrEventNotFound).SetFunction("UpdateEvent")
	}

	recordUniqueDocument := validationsDomain.ValidateUniqueFieldParams{
		Table:            "core_events",
		ColumnName:       "document",
		Value:            body.Document,
		RecordIdName:     "id",
		RecordIdValue:    eventId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var uniqueDocument bool
	uniqueDocument, err = u.validationRepository.ValidateUniqueField(ctx, recordUniqueDocument)
	if err != nil {
		return err
	}
	if !uniqueDocument {
		return eventsDomain.ErrEventDocumentAlreadyExist.SetFunction("UpdateEvent")
	}

	recordUniqueParams := validationsDomain.ValidateUniqueFieldParams{
		Table:            "core_events",
		ColumnName:       "code",
		Value:            body.Code,
		RecordIdName:     "id",
		RecordIdValue:    eventId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	unique, err := u.validationRepository.ValidateUniqueField(ctx, recordUniqueParams)
	if err != nil {
		return err
	}
	if !unique {
		return u.err.Clone().CopyCodeDescription(eventsDomain.ErrEventCodeAlreadyExist).
			SetFunction("UpdateEvent").SetLayer(logErrorCoreDomain.UseCase)
	}

	// review roles by default
	//defaultToEvent, err := u.roleDefaultsRepository.GetRolesDefaults(ctx, &eventId, nil)
	//if err != nil {
	//	return err
	//}
	err = u.eventsRepository.UpdateEvent(ctx, eventId, body)
	return
}

func (u eventsUseCase) DeleteEvent(
	ctx context.Context,
	eventId string,
) (
	update bool,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	var deleted string
	deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_events",
		IdColumnName:     "id",
		IdValue:          eventId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, eventsDomain.ErrEventIdHasBeenDeleted
	}

	res, err := u.eventsRepository.DeleteEvent(ctx, eventId)
	return res, err
}

func (u eventsUseCase) GetRolesByEvent(
	ctx context.Context,
	eventId string,
) (
	res []eventsDomain.Role,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err = u.eventsRepository.GetRolesByEvent(ctx, eventId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u eventsUseCase) EnableDisableEvent(
	ctx context.Context,
	eventId string,
	body eventsDomain.EnableDisableEventRequest,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "core_events",
		IdColumnName:     "id",
		IdValue:          eventId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(eventsDomain.ErrEventNotFound).SetFunction("EnableDisableEvent")
	}

	err = u.eventsRepository.EnableDisableEvent(ctx, eventId, body)
	return
}

func (u eventsUseCase) GetEventSummary(
	ctx context.Context,
	params eventsDomain.GetEventSumsParams,
) (
	res []eventsDomain.EventSums,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetEventSums error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetEventSums, &wg)
		res, errGetEventSums = u.eventsRepository.GetEventSums(ctx, params)
		wg.Done()
	}()
	wg.Wait()

	if errGetEventSums != nil {
		return nil, errGetEventSums
	}

	return res, nil
}
