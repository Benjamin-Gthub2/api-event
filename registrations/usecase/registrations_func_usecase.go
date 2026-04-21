/*
 * File: warehouse_transfers_receipts_func_usecase.go
 * Author: Benjamin
 * Copyright: 2025, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This contains the use case functions.
 *
 * Last Modified: 2024-02-01
 */

package usecase

import (
	"context"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

func (u registrationsUseCase) GetQrRegistrationById(
	ctx context.Context,
	registrationId string,
) (
	qrCode []byte,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "registrations",
		IdColumnName:     "id",
		IdValue:          registrationId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return qrCode, err
	}
	if !exist {
		return qrCode, registrationsDomain.ErrRegistrationsNotFound
	}

	qrCode, err = u.registrationsRepository.GetQrRegistrationById(ctx, registrationId)
	if err != nil {
		return qrCode, err
	}
	return qrCode, nil
}

func (u registrationsUseCase) GetRegistrationById(
	ctx context.Context,
	registrationId string,
) (
	requirement *registrationsDomain.Registration,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var deleted = "deleted_at"
	recordExistsParams := validationsDomain.RecordExistsParams{
		Table:            "registrations",
		IdColumnName:     "id",
		IdValue:          registrationId,
		StatusColumnName: &deleted,
		StatusValue:      nil,
	}
	var exist bool
	exist, err = u.validationRepository.RecordExists(ctx, recordExistsParams)
	if err != nil {
		return requirement, err
	}
	if !exist {
		return requirement, registrationsDomain.ErrUseCaseRegistrationsNotFound
	}

	requirement, err = u.registrationsRepository.GetRegistrationById(ctx, registrationId)
	if err != nil {
		return requirement, err
	}
	return requirement, nil
}
