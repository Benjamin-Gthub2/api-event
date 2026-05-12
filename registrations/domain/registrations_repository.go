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
	"database/sql"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type RegistrationsRepository interface {
	GetRegistrationById(ctx context.Context, registrationId string) (*Registration, error)
	GetRegistrations(
		ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetRegistrationsParams,
	) ([]Registration, error)
	GetTotalRegistrations(
		ctx context.Context, searchParams GetRegistrationsParams) (*int, error)
	CreateRegistration(ctx context.Context, tx *sql.Tx, body CreateRegistration) error
	MainCreateRegistration(ctx context.Context, body CreateRegistration) (err error)
	UpdateRegistrationStatus(ctx context.Context, registrationId string, statusCode string) error
	UpdateRegistrationSendQr(ctx context.Context, registrationId string) error
}
