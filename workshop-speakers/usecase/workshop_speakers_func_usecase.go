package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

func (u workshopSpeakersUseCase) GetWorkshopSpeakerById(
	ctx context.Context,
	workshopSpeakerId string,
) (
	workshopSpeaker *workshopSpeakersDomain.WorkshopSpeaker,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "workshop_speakers",
		IdColumnName:     "id",
		IdValue:          workshopSpeakerId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, workshopSpeakersDomain.ErrWorkshopSpeakerNotFound
	}

	workshopSpeaker, err = u.workshopSpeakersRepository.GetWorkshopSpeakerById(ctx, workshopSpeakerId)
	if err != nil {
		return nil, err
	}
	return workshopSpeaker, nil
}

func (u workshopSpeakersUseCase) GetWorkshopSpeakers(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams workshopSpeakersDomain.GetWorkshopSpeakersParams,
) (
	res []workshopSpeakersDomain.WorkshopSpeaker,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetWorkshopSpeakers, errGetTotal error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetWorkshopSpeakers, &wg)
		res, errGetWorkshopSpeakers = u.workshopSpeakersRepository.GetWorkshopSpeakers(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotal, &wg)
		total, errGetTotal = u.workshopSpeakersRepository.GetTotalWorkshopSpeakers(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetWorkshopSpeakers != nil {
		return nil, nil, errGetWorkshopSpeakers
	}
	if errGetTotal != nil {
		return nil, nil, errGetTotal
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u workshopSpeakersUseCase) CreateWorkshopSpeaker(
	ctx context.Context,
	userId string,
	body workshopSpeakersDomain.CreateWorkshopSpeakerBody,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	workshopSpeakerId := uuid.New().String()

	createWorkshopSpeaker := workshopSpeakersDomain.CreateWorkshopSpeaker{
		Id:         workshopSpeakerId,
		WorkshopId: body.WorkshopId,
		SpeakerId:  body.SpeakerId,
		CreatedBy:  userId,
	}
	err = u.workshopSpeakersRepository.CreateWorkshopSpeaker(ctx, createWorkshopSpeaker)
	if err != nil {
		return nil, err
	}
	return &workshopSpeakerId, nil
}

func (u workshopSpeakersUseCase) DeleteWorkshopSpeaker(
	ctx context.Context,
	workshopSpeakerId string,
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
		Table:            "workshop_speakers",
		IdColumnName:     "id",
		IdValue:          workshopSpeakerId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return workshopSpeakersDomain.ErrWorkshopSpeakerNotFound
	}

	deleteWorkshopSpeaker := workshopSpeakersDomain.DeleteWorkshopSpeaker{
		Id:        workshopSpeakerId,
		DeletedBy: userId,
	}
	err = u.workshopSpeakersRepository.DeleteWorkshopSpeaker(ctx, deleteWorkshopSpeaker)
	if err != nil {
		return err
	}
	return nil
}
