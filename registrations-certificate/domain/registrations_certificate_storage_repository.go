/*
 * File: registrations_certificate_storage_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the storage repository interface for the registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package domain

import "context"

type RegistrationCertificateStorageRepository interface {
	UploadCertificate(ctx context.Context, registrationId string, data []byte) error
	// GetCertificate returns (nil, nil) when the certificate does not exist yet.
	GetCertificate(ctx context.Context, registrationId string) ([]byte, error)
}
