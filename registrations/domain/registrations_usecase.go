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

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type RegistrationsUseCase interface {
	GetQrRegistrationById(ctx context.Context, registrationId string) ([]byte, error)
	GetRegistrationById(ctx context.Context, registrationId string) (*Registration, error)
	GetRegistrations(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetRegistrationsParams,
	) ([]Registration, *paramsDomain.PaginationResults, error)
	CreateRegistration(ctx context.Context, userId string, body CreateRegistrationBody) (*string, error)
	UpdateRegistrationStatus(ctx context.Context, registrationId string, statusCode string) error
}
