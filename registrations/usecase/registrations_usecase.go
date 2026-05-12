/*
 * File: registrations_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the use case of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package usecase

import (
	"time"

	authDomain "github.com/Benjamin-Gthub2/api-shared/auth/domain"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	validationsDomain "github.com/Benjamin-Gthub2/api-shared/validations/domain"

	registrationSharedDomain "github.com/Benjamin-Gthub2/api-event/registrations-shared/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsUseCase struct {
	registrationsRepository         registrationsDomain.RegistrationsRepository
	registrationsRTRepository       registrationsDomain.RegistrationsRTRepository
	registrationsWhatsAppRepository registrationsDomain.RegistrationsWhatsAppRepository
	registrationSharedRepository    registrationSharedDomain.RegistrationSharedRepository
	validationRepository            validationsDomain.ValidationRepository
	authRepository                  authDomain.AuthRepository
	contextTimeout                  time.Duration
	err                             *errDomain.SmartError
}

func NewRegistrationsUseCase(
	ur registrationsDomain.RegistrationsRepository,
	registrationsRTRepository registrationsDomain.RegistrationsRTRepository,
	registrationsWhatsAppRepository registrationsDomain.RegistrationsWhatsAppRepository,
	registrationSharedRepository registrationSharedDomain.RegistrationSharedRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) registrationsDomain.RegistrationsUseCase {
	return &registrationsUseCase{
		registrationsRepository:         ur,
		registrationsRTRepository:       registrationsRTRepository,
		registrationsWhatsAppRepository:  registrationsWhatsAppRepository,
		registrationSharedRepository:    registrationSharedRepository,
		validationRepository:            validation,
		authRepository:                  authRepository,
		contextTimeout:                  timeout,
		err:                             errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
