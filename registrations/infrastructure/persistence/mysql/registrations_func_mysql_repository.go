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
	_ "embed"

	"github.com/skip2/go-qrcode"

	logErrorCoreDomain "github.com/smart0n3/api-shared/error-core/domain"
)

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
