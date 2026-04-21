/*
 * File: registrations_receipts_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package domain

import (
	"context"
)

type RegistrationsUseCase interface {
	GetQrRegistrationById(ctx context.Context, registrationId string) ([]byte, error)
}
