/*
 * File: registrations_certificate_pdf_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository constructor for registrations certificate PDF generation.
 *
 * Last Modified: 2026-05-16
 */

package registrations

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationCertificateDomain "github.com/Benjamin-Gthub2/api-event/registrations-certificate/domain"
)

// pdfSemaphore limits concurrent wkhtmltopdf processes to 1 to avoid OOM on low-memory hosts.
var pdfSemaphore = make(chan struct{}, 1)

type registrationCertificatesReportPdfRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRegistrationCertificatePdfRepository(
	clock smartClock.Clock,
	timeout int,
) registrationCertificateDomain.RegistrationCertificateRepository {
	return &registrationCertificatesReportPdfRepo{
		clock:   clock,
		timeout: time.Duration(timeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
}
