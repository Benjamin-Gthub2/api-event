/*
 * File: registrations_storage_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the storage repository interface for registrations.
 *
 * Last Modified: 2026-05-12
 */

package domain

import "context"

type RegistrationsStorageRepository interface {
	UploadQr(ctx context.Context, registrationId string, data []byte) error
	GetQr(ctx context.Context, registrationId string) ([]byte, error)
}
