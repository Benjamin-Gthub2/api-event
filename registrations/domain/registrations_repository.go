/*
 * File: registrations_repository.go
 * Author: benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository in the domain.
 *
 * Last Modified: 2026-04-21
 */

package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type RegistrationsRepository interface {
	GetQrRegistrationById(ctx context.Context, registrationId string) ([]byte, error)
	GetRegistrationById(ctx context.Context, registrationId string) (*Registration, error)
	GetRegistrations(
		ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetRegistrationsParams,
	) ([]Registration, error)
	GetTotalRegistrations(
		ctx context.Context, searchParams GetRegistrationsParams) (*int, error)
}
