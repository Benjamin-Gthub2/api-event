/*
 * File: registrations_certificate_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case struct and constructor for registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package usecase

import (
	"time"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationCertificateDomain "github.com/Benjamin-Gthub2/api-event/registrations-certificate/domain"
	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsCertificateUseCase struct {
	registrationsRepository              registrationsDomain.RegistrationsRepository
	registrationCertificateRepository    registrationCertificateDomain.RegistrationCertificateRepository
	registrationCertificateStorageRepository registrationCertificateDomain.RegistrationCertificateStorageRepository
	contextTimeout                       time.Duration
	err                                  *errDomain.SmartError
}

func NewRegistrationsCertificateUseCase(
	registrationsRepository registrationsDomain.RegistrationsRepository,
	registrationCertificateRepository registrationCertificateDomain.RegistrationCertificateRepository,
	registrationCertificateStorageRepository registrationCertificateDomain.RegistrationCertificateStorageRepository,
	timeout time.Duration,
) registrationCertificateDomain.RegistrationsCertificateUseCase {
	return &registrationsCertificateUseCase{
		registrationsRepository:              registrationsRepository,
		registrationCertificateRepository:    registrationCertificateRepository,
		registrationCertificateStorageRepository: registrationCertificateStorageRepository,
		contextTimeout:                       timeout,
		err:                                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
