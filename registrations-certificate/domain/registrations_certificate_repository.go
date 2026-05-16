/*
 * File: registrations_certificate_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository interface for the registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package domain

import (
	"context"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type RegistrationCertificateRepository interface {
	GenerateRegistrationCertificatePdf(ctx context.Context, registration *registrationsDomain.Registration) ([]byte, error)
}
