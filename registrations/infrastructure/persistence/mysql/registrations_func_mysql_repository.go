/*
 * File: registrations_func_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository functions of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package mysql

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/jackskj/carta"
	"github.com/skip2/go-qrcode"
	"github.com/smart0n3/api-shared/db"
	paramsDomain "github.com/smart0n3/api-shared/params/domain"
	"github.com/stroiman/go-automapper"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

//go:embed sql/get_registration_by_id.sql
var QueryGetRegistrationById string

//go:embed sql/get_registrations.sql
var QueryGetRegistrations string

//go:embed sql/get_total_registrations.sql
var QueryGetTotalRegistrations string

//go:embed sql/create_registration.sql
var QueryCreateRegistration string

func (r registrationsMySQLRepo) GetQrRegistrationById(
	ctx context.Context,
	registrationId string,
) (
	qrCode []byte,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, r.timeout)
	defer cancel()

	qrCode, err = qrcode.Encode(registrationId, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

func (r registrationsMySQLRepo) GetRegistrationById(
	ctx context.Context,
	registrationId string,
) (
	registrationById *registrationsDomain.Registration,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}
	results, err := client.
		QueryContext(
			ctx,
			QueryGetRegistrationById,
			registrationId,
		)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}
	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}(results)

	RegistrationByIdTmp := make([]Registration, 0)
	err = carta.Map(results, &RegistrationByIdTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrationById").SetRaw(err)
	}

	registrationAux := make([]registrationsDomain.Registration, 0)
	automapper.Map(RegistrationByIdTmp, &registrationAux)
	if len(registrationAux) == 0 {
		return nil, registrationsDomain.ErrRegistrationsNotFound
	}

	return &registrationAux[0], nil
}

func (r registrationsMySQLRepo) GetRegistrations(
	ctx context.Context,
	pagination paramsDomain.PaginationParams,
	searchParams registrationsDomain.GetRegistrationsParams,
) (
	registrationsRows []registrationsDomain.Registration,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	sizePage := pagination.GetSizePage()
	offset := pagination.GetOffset()

	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}

	results, err := client.QueryContext(
		ctx,
		QueryGetRegistrations,
		searchParams.StartDate,
		searchParams.StartDate,
		searchParams.EndDate,
		searchParams.CreatedBy,
		searchParams.CreatedBy,
		sizePage,
		offset,
	)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}

	defer func(results *sql.Rows) {
		errClose := results.Close()
		if errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &err)
		}
	}(results)

	registrationsTmp := make([]Registration, 0)
	err = carta.Map(results, &registrationsTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetRegistrations").SetRaw(err)
	}
	var registrations = make([]registrationsDomain.Registration, 0)
	automapper.Map(registrationsTmp, &registrations)

	return registrations, nil
}

func (r registrationsMySQLRepo) GetTotalRegistrations(
	ctx context.Context,
	searchParams registrationsDomain.GetRegistrationsParams,
) (
	total *int,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	var totalTmp int
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrations").SetRaw(err)
	}
	err = client.
		QueryRowContext(
			ctx,
			QueryGetTotalRegistrations,
			searchParams.StartDate,
			searchParams.StartDate,
			searchParams.EndDate,
			searchParams.CreatedBy,
			searchParams.CreatedBy,
		).Scan(&totalTmp)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetTotalRegistrations").SetRaw(err)
	}
	total = &totalTmp
	return total, nil
}

func (r registrationsMySQLRepo) CreateRegistration(
	ctx context.Context,
	body registrationsDomain.CreateRegistration,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)
	now := r.clock.Now().Format("2006-01-02 15:04:06")
	client, _, err := db.ClientDB(ctx)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRegistration").SetRaw(err)
	}
	_, err = client.ExecContext(ctx,
		QueryCreateRegistration,
		body.Id,
		body.SessionId,
		body.BeneficiaryId,
		body.CreatedBy,
		now,
	)
	if err != nil {
		return r.err.Clone().SetFunction("CreateRegistration").SetRaw(err)
	}
	return
}
