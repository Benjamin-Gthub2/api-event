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
)

type RegistrationsRepository interface {
	GetQrRegistrationById(ctx context.Context, registrationId string) ([]byte, error)
	GetRegistrationById(ctx context.Context, registrationId string) (*Registration, error)
}
