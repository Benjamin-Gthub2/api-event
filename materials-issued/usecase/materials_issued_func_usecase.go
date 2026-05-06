package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

func (u materialsIssuedUseCase) GetMaterialIssuedById(
	ctx context.Context,
	materialIssuedId string,
) (
	materialIssued *materialsIssuedDomain.MaterialIssued,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "materials_issued",
		IdColumnName:     "id",
		IdValue:          materialIssuedId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, materialsIssuedDomain.ErrMaterialIssuedNotFound
	}

	materialIssued, err = u.materialsIssuedRepository.GetMaterialIssuedById(ctx, materialIssuedId)
	if err != nil {
		return nil, err
	}
	return materialIssued, nil
}

func (u materialsIssuedUseCase) GetMaterialsIssued(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams materialsIssuedDomain.GetMaterialsIssuedParams,
) (
	res []materialsIssuedDomain.MaterialIssued,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetMaterials, errGetTotal error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetMaterials, &wg)
		res, errGetMaterials = u.materialsIssuedRepository.GetMaterialsIssued(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotal, &wg)
		total, errGetTotal = u.materialsIssuedRepository.GetTotalMaterialsIssued(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetMaterials != nil {
		return nil, nil, errGetMaterials
	}
	if errGetTotal != nil {
		return nil, nil, errGetTotal
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u materialsIssuedUseCase) CreateMaterialIssued(
	ctx context.Context,
	userId string,
	body materialsIssuedDomain.CreateMaterialIssuedBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	materialIssuedId := uuid.New().String()
	createMaterialIssued := materialsIssuedDomain.CreateMaterialIssued{
		Id:          materialIssuedId,
		Description: body.Description,
		CreatedBy:   userId,
	}
	err = u.materialsIssuedRepository.CreateMaterialIssued(ctx, createMaterialIssued)
	if err != nil {
		return nil, err
	}
	return &materialIssuedId, nil
}

func (u materialsIssuedUseCase) UpdateMaterialIssued(
	ctx context.Context,
	materialIssuedId string,
	body materialsIssuedDomain.UpdateMaterialIssuedBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "materials_issued",
		IdColumnName:     "id",
		IdValue:          materialIssuedId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return materialsIssuedDomain.ErrMaterialIssuedNotFound
	}

	updateMaterialIssued := materialsIssuedDomain.UpdateMaterialIssued{
		Id:          materialIssuedId,
		Description: body.Description,
	}
	err = u.materialsIssuedRepository.UpdateMaterialIssued(ctx, updateMaterialIssued)
	if err != nil {
		return err
	}
	return nil
}

func (u materialsIssuedUseCase) DeleteMaterialIssued(
	ctx context.Context,
	materialIssuedId string,
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
		Table:            "materials_issued",
		IdColumnName:     "id",
		IdValue:          materialIssuedId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return materialsIssuedDomain.ErrMaterialIssuedNotFound
	}

	deleteMaterialIssued := materialsIssuedDomain.DeleteMaterialIssued{
		Id:        materialIssuedId,
		DeletedBy: userId,
	}
	err = u.materialsIssuedRepository.DeleteMaterialIssued(ctx, deleteMaterialIssued)
	if err != nil {
		return err
	}
	return nil
}
