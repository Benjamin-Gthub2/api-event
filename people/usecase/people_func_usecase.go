/*
 * File: people_func_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

func (u peopleUseCase) GetPeople(
	ctx context.Context,
	searchParams domain.GetPeopleParams,
	pagination paramsDomain.PaginationParams,
) (
	people []domain.People,
	resultPagination *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetPeople, errGetTotalPeople error
	var total *int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		people, errGetPeople = u.peopleRepository.GetPeople(ctx, searchParams, pagination)
		wg.Done()
	}()
	go func() {
		total, errGetTotalPeople = u.peopleRepository.GetTotalPeople(ctx, searchParams, pagination)
		wg.Done()
	}()
	wg.Wait()

	if errGetPeople != nil {
		return nil, nil, errGetPeople
	}
	if errGetTotalPeople != nil {
		return nil, nil, errGetTotalPeople
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return people, &paginationRes, nil
}

func (u peopleUseCase) CreatePerson(
	ctx context.Context,
	body domain.CreatePersonBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "people",
		IdColumnName:     "user_id",
		IdValue:          body.UserId,
		StatusColumnName: nil,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.ValidateExistence(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, domain.ErrPersonUserIdAlreadyExist
	}
	peopleId := uuid.New().String()
	id, err = u.peopleRepository.CreatePerson(ctx, peopleId, body)
	return
}

func (u peopleUseCase) UpdatePerson(
	ctx context.Context,
	peopleId string,
	body domain.UpdatePersonBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:        "people",
		IdColumnName: "id",
		IdValue:      peopleId,
	}
	exist, err := u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return u.err.Clone().CopyCodeDescription(domain.ErrPersonNotFound).SetFunction("UpdatePeople")
	}

	err = u.peopleRepository.UpdatePerson(ctx, peopleId, body)
	return
}

func (u peopleUseCase) DeletePerson(
	ctx context.Context,
	peopleId string,
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
		Table:            "people",
		IdColumnName:     "id",
		IdValue:          peopleId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, domain.ErrPersonIdAlreadyDeleted
	}
	update, err = u.peopleRepository.DeletePerson(ctx, peopleId)
	return update, nil

}
