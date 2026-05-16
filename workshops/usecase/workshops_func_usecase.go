package usecase

import (
	"context"
	"sync"
	"time"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"
	"github.com/google/uuid"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

var limaLoc = time.FixedZone("America/Lima", -5*60*60)

func (u workshopsUseCase) GetWorkshopById(
	ctx context.Context,
	workshopId string,
) (
	workshop *workshopsDomain.Workshop,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "workshops",
		IdColumnName:     "id",
		IdValue:          workshopId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, workshopsDomain.ErrWorkshopNotFound
	}

	workshop, err = u.workshopsRepository.GetWorkshopById(ctx, workshopId)
	if err != nil {
		return nil, err
	}
	return workshop, nil
}

func (u workshopsUseCase) GetWorkshops(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopsDomain.GetWorkshopsParams,
) (
	res []workshopsDomain.Workshop,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if searchParams.OnlyToday != nil && *searchParams.OnlyToday {
		today := time.Now().In(limaLoc).Format("2006-01-02")
		searchParams.StartDate = &today
	}

	var errGetWorkshops, errGetTotalWorkshops error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetWorkshops, &wg)
		res, errGetWorkshops = u.workshopsRepository.GetWorkshops(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotalWorkshops, &wg)
		total, errGetTotalWorkshops = u.workshopsRepository.GetTotalWorkshops(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetWorkshops != nil {
		return nil, nil, errGetWorkshops
	}
	if errGetTotalWorkshops != nil {
		return nil, nil, errGetTotalWorkshops
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u workshopsUseCase) CreateWorkshop(
	ctx context.Context,
	userId string,
	body workshopsDomain.CreateWorkshopBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	workshopId := uuid.New().String()

	createWorkshop := workshopsDomain.CreateWorkshop{
		Id:        workshopId,
		TypeId:    body.TypeId,
		Name:      body.Name,
		Shortname: body.Shortname,
		Code:      body.Code,
		Capacity:  body.Capacity,
		StartDate: body.StartDate.In(limaLoc).Format("2006-01-02 15:04:05"),
		EndDate:   body.EndDate.In(limaLoc).Format("2006-01-02 15:04:05"),
		Place:     body.Place,
		EventId:   body.EventId,
		CreatedBy: userId,
	}
	err = u.workshopsRepository.CreateWorkshop(ctx, createWorkshop)
	if err != nil {
		return nil, err
	}
	return &workshopId, nil
}

func (u workshopsUseCase) UpdateWorkshop(
	ctx context.Context,
	workshopId string,
	body workshopsDomain.UpdateWorkshopBody,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "workshops",
		IdColumnName:     "id",
		IdValue:          workshopId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return workshopsDomain.ErrWorkshopNotFound
	}

	updateWorkshop := workshopsDomain.UpdateWorkshop{
		Id:        workshopId,
		TypeId:    body.TypeId,
		Name:      body.Name,
		Shortname: body.Shortname,
		Code:      body.Code,
		Capacity:  body.Capacity,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
		Place:     body.Place,
	}
	err = u.workshopsRepository.UpdateWorkshop(ctx, updateWorkshop)
	if err != nil {
		return err
	}
	return nil
}

func (u workshopsUseCase) DeleteWorkshop(
	ctx context.Context,
	workshopId string,
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
		Table:            "workshops",
		IdColumnName:     "id",
		IdValue:          workshopId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return workshopsDomain.ErrWorkshopNotFound
	}

	deleteWorkshop := workshopsDomain.DeleteWorkshop{
		Id:        workshopId,
		DeletedBy: userId,
	}
	err = u.workshopsRepository.DeleteWorkshop(ctx, deleteWorkshop)
	if err != nil {
		return err
	}
	return nil
}

func (u workshopsUseCase) GetWorkshopSummary(
	ctx context.Context,
	searchParams workshopsDomain.GetWorkshopSumsParams,
) (
	res []workshopsDomain.WorkshopSums,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetWorkshopSums error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetWorkshopSums, &wg)
		res, errGetWorkshopSums = u.workshopsRepository.GetWorkshopSums(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetWorkshopSums != nil {
		return nil, errGetWorkshopSums
	}

	return res, nil
}
