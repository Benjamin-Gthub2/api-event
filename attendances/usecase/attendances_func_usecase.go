package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

func (u attendancesUseCase) GetAttendanceById(
	ctx context.Context,
	attendanceId string,
) (
	attendance *attendancesDomain.Attendance,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "attendances",
		IdColumnName:     "id",
		IdValue:          attendanceId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, attendancesDomain.ErrAttendanceNotFound
	}

	attendance, err = u.attendancesRepository.GetAttendanceById(ctx, attendanceId)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (u attendancesUseCase) GetAttendances(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams attendancesDomain.GetAttendancesParams,
) (
	res []attendancesDomain.Attendance,
	paginationResult *paramsDomain.PaginationResults,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var errGetAttendances, errGetTotal error
	var total *int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetAttendances, &wg)
		res, errGetAttendances = u.attendancesRepository.GetAttendances(ctx, pagination, searchParams)
		wg.Done()
	}()
	go func() {
		defer logErrorCoreDomain.PanicThreadRecovery(&ctx, &errGetTotal, &wg)
		total, errGetTotal = u.attendancesRepository.GetTotalAttendances(ctx, searchParams)
		wg.Done()
	}()
	wg.Wait()

	if errGetAttendances != nil {
		return nil, nil, errGetAttendances
	}
	if errGetTotal != nil {
		return nil, nil, errGetTotal
	}

	paginationRes := paramsDomain.PaginationResults{}
	paginationRes.FromParams(pagination, *total)

	return res, &paginationRes, nil
}

func (u attendancesUseCase) CreateAttendance(
	ctx context.Context,
	userId string,
) (
	id *string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	attendanceId := uuid.New().String()
	createAttendance := attendancesDomain.CreateAttendance{
		Id:        attendanceId,
		CreatedBy: userId,
	}
	err = u.attendancesRepository.CreateAttendance(ctx, createAttendance)
	if err != nil {
		return nil, err
	}
	return &attendanceId, nil
}

func (u attendancesUseCase) DeleteAttendance(
	ctx context.Context,
	attendanceId string,
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
		Table:            "attendances",
		IdColumnName:     "id",
		IdValue:          attendanceId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return err
	}
	if !exist {
		return attendancesDomain.ErrAttendanceNotFound
	}

	deleteAttendance := attendancesDomain.DeleteAttendance{
		Id:        attendanceId,
		DeletedBy: userId,
	}
	err = u.attendancesRepository.DeleteAttendance(ctx, deleteAttendance)
	if err != nil {
		return err
	}
	return nil
}
