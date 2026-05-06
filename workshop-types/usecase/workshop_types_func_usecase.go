package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

func (u workshopTypesUseCase) GetWorkshopTypeById(
	ctx context.Context,
	workshopTypeId string,
) (
	workshopType *workshopTypesDomain.WorkshopType,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "workshop_types",
		IdColumnName:     "id",
		IdValue:          workshopTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, workshopTypesDomain.ErrWorkshopTypeNotFound
	}

	workshopType, err = u.workshopTypesRepository.GetWorkshopTypeById(ctx, workshopTypeId)
	if err != nil {
		return nil, err
	}
	return workshopType, nil
}

func (u workshopTypesUseCase) GetWorkshopTypes(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopTypesDomain.GetWorkshopTypesParams,
) (
	res []workshopTypesDomain.WorkshopType,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetWorkshopTypes, errGetTotalWorkshopTypes error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetWorkshopTypes, &wg)
		res, errGetWorkshopTypes = u.workshopTypesRepository.GetWorkshopTypes(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalWorkshopTypes, &wg)
		total, errGetTotalWorkshopTypes = u.workshopTypesRepository.GetTotalWorkshopTypes(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetWorkshopTypes != nil {
		return nil, nil, errGetWorkshopTypes
	}
	if errGetTotalWorkshopTypes != nil {
		return nil, nil, errGetTotalWorkshopTypes
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u workshopTypesUseCase) CreateWorkshopType(
	ctx context.Context,
	userId string,
	body workshopTypesDomain.CreateWorkshopTypeBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	workshopTypeId := uuid.New().String()

	createWorkshopType := workshopTypesDomain.CreateWorkshopType{
		Id:          workshopTypeId,
		Code:        body.Code,
		Description: body.Description,
		Enable:      body.Enable,
		CreatedBy:   userId,
	}
	err = u.workshopTypesRepository.CreateWorkshopType(ctx, createWorkshopType)
	if err != nil {
		return nil, err
	}
	return &workshopTypeId, nil
}

func (u workshopTypesUseCase) UpdateWorkshopType(
	ctx context.Context,
	workshopTypeId string,
	body workshopTypesDomain.UpdateWorkshopTypeBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "workshop_types",
		IdColumnName:     "id",
		IdValue:          workshopTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return workshopTypesDomain.ErrWorkshopTypeNotFound
	}

	updateWorkshopType := workshopTypesDomain.UpdateWorkshopType{
		Id:          workshopTypeId,
		Code:        body.Code,
		Description: body.Description,
		Enable:      body.Enable,
	}
	err = u.workshopTypesRepository.UpdateWorkshopType(ctx, updateWorkshopType)
	if err != nil {
		return err
	}
	return nil
}

func (u workshopTypesUseCase) DeleteWorkshopType(
	ctx context.Context,
	workshopTypeId string,
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
		Table:            "workshop_types",
		IdColumnName:     "id",
		IdValue:          workshopTypeId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return workshopTypesDomain.ErrWorkshopTypeNotFound
	}

	deleteWorkshopType := workshopTypesDomain.DeleteWorkshopType{
		Id:        workshopTypeId,
		DeletedBy: userId,
	}
	err = u.workshopTypesRepository.DeleteWorkshopType(ctx, deleteWorkshopType)
	if err != nil {
		return err
	}
	return nil
}
