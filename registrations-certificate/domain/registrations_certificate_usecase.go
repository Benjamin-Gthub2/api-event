/*
 * File: registrations_certificate_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case interface for the registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package domain

import (
	"context"
)

type RegistrationsCertificateUseCase interface {
	GenerateRegistrationsCertificatePdf(ctx context.Context, registrationId string) (pdfBytes []byte, fileName string, err error)
}
