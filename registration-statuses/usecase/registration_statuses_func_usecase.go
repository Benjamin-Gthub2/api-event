package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

func (u registrationStatusesUseCase) GetRegistrationStatusById(
	ctx context.Context,
	registrationStatusId string,
) (
	registrationStatus *registrationStatusesDomain.RegistrationStatus,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "registration_statuses",
		IdColumnName:     "id",
		IdValue:          registrationStatusId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, registrationStatusesDomain.ErrRegistrationStatusNotFound
	}

	registrationStatus, err = u.registrationStatusesRepository.GetRegistrationStatusById(ctx, registrationStatusId)
	if err != nil {
		return nil, err
	}
	return registrationStatus, nil
}

func (u registrationStatusesUseCase) GetRegistrationStatuses(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams registrationStatusesDomain.GetRegistrationStatusesParams,
) (
	res []registrationStatusesDomain.RegistrationStatus,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetStatuses, errGetTotal error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetStatuses, &wg)
		res, errGetStatuses = u.registrationStatusesRepository.GetRegistrationStatuses(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotal, &wg)
		total, errGetTotal = u.registrationStatusesRepository.GetTotalRegistrationStatuses(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetStatuses != nil {
		return nil, nil, errGetStatuses
	}
	if errGetTotal != nil {
		return nil, nil, errGetTotal
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u registrationStatusesUseCase) CreateRegistrationStatus(
	ctx context.Context,
	body registrationStatusesDomain.CreateRegistrationStatusBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	registrationStatusId := uuid.New().String()
	createRegistrationStatus := registrationStatusesDomain.CreateRegistrationStatus{
		Id:          registrationStatusId,
		Code:        body.Code,
		Description: body.Description,
		Position:    body.Position,
		Enable:      body.Enable,
	}
	err = u.registrationStatusesRepository.CreateRegistrationStatus(ctx, createRegistrationStatus)
	if err != nil {
		return nil, err
	}
	return &registrationStatusId, nil
}

func (u registrationStatusesUseCase) UpdateRegistrationStatus(
	ctx context.Context,
	registrationStatusId string,
	body registrationStatusesDomain.UpdateRegistrationStatusBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "registration_statuses",
		IdColumnName:     "id",
		IdValue:          registrationStatusId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return registrationStatusesDomain.ErrRegistrationStatusNotFound
	}

	updateRegistrationStatus := registrationStatusesDomain.UpdateRegistrationStatus{
		Id:          registrationStatusId,
		Code:        body.Code,
		Description: body.Description,
		Position:    body.Position,
		Enable:      body.Enable,
	}
	err = u.registrationStatusesRepository.UpdateRegistrationStatus(ctx, updateRegistrationStatus)
	if err != nil {
		return err
	}
	return nil
}

func (u registrationStatusesUseCase) DeleteRegistrationStatus(
	ctx context.Context,
	registrationStatusId string,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "registration_statuses",
		IdColumnName:     "id",
		IdValue:          registrationStatusId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return registrationStatusesDomain.ErrRegistrationStatusNotFound
	}

	deleteRegistrationStatus := registrationStatusesDomain.DeleteRegistrationStatus{
		Id: registrationStatusId,
	}
	err = u.registrationStatusesRepository.DeleteRegistrationStatus(ctx, deleteRegistrationStatus)
	if err != nil {
		return err
	}
	return nil
}
