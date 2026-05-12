/*
 * File: registrations_certificate_pdf_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository for registrations certificate.
 *
 * Last Modified: 2026-05-12
 */

package registrations

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"

	registrationCertificateDomain "github.com/Benjamin-Gthub2/api-event/registrations-certificate/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

type registrationCertificatesReportPdfRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRequirementsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) registrationCertificateDomain.RegistrationCertificateRepository {
	rep := &registrationCertificatesReportPdfRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
