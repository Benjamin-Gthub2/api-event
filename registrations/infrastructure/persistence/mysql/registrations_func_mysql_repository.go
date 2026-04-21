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
	"github.com/stroiman/go-automapper"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

//go:embed sql/get_registration_by_id.sql
var QueryGetRegistrationById string

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
