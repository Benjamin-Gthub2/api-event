package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

func (u eventTypesUseCase) GetEventTypeById(
	ctx context.Context,
	eventTypeId string,
) (
	eventType *eventTypesDomain.EventType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "event_types",
		IdColumnName:     "id",
		IdValue:          eventTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, eventTypesDomain.ErrEventTypeNotFound
	}

	eventType, err = u.eventTypesRepository.GetEventTypeById(ctx, eventTypeId)
	if err != nil {
		return nil, err
	}
	return eventType, nil
}

func (u eventTypesUseCase) GetEventTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams eventTypesDomain.GetEventTypesParams,
) (
	res []eventTypesDomain.EventType,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetEventTypes, errGetTotalEventTypes error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetEventTypes, &wg)
		res, errGetEventTypes = u.eventTypesRepository.GetEventTypes(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalEventTypes, &wg)
		total, errGetTotalEventTypes = u.eventTypesRepository.GetTotalEventTypes(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetEventTypes != nil {
		return nil, nil, errGetEventTypes
	}
	if errGetTotalEventTypes != nil {
		return nil, nil, errGetTotalEventTypes
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u eventTypesUseCase) CreateEventType(
	ctx context.Context,
	userId string,
	body eventTypesDomain.CreateEventTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	eventTypeId := uuid.New().String()

	createEventType := eventTypesDomain.CreateEventType{
		Id:          eventTypeId,
		Code:        body.Code,
		Description: body.Description,
		Enable:      body.Enable,
		CreatedBy:   userId,
	}
	err = u.eventTypesRepository.CreateEventType(ctx, createEventType)
	if err != nil {
		return nil, err
	}
	return &eventTypeId, nil
}

func (u eventTypesUseCase) UpdateEventType(
	ctx context.Context,
	eventTypeId string,
	body eventTypesDomain.UpdateEventTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "event_types",
		IdColumnName:     "id",
		IdValue:          eventTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return eventTypesDomain.ErrEventTypeNotFound
	}

	updateEventType := eventTypesDomain.UpdateEventType{
		Id:          eventTypeId,
		Code:        body.Code,
		Description: body.Description,
		Enable:      body.Enable,
	}
	err = u.eventTypesRepository.UpdateEventType(ctx, updateEventType)
	if err != nil {
		return err
	}
	return nil
}

func (u eventTypesUseCase) DeleteEventType(
	ctx context.Context,
	eventTypeId string,
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
		Table:            "event_types",
		IdColumnName:     "id",
		IdValue:          eventTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return eventTypesDomain.ErrEventTypeNotFound
	}

	deleteEventType := eventTypesDomain.DeleteEventType{
		Id:        eventTypeId,
		DeletedBy: userId,
	}
	err = u.eventTypesRepository.DeleteEventType(ctx, deleteEventType)
	if err != nil {
		return err
	}
	return nil
}
