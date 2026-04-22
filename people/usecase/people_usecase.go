/*
 * File: people_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package usecase

import (
	"time"

	authDomain "github.com/smart0n3/api-shared/auth/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	"github.com/Benjamin-Gthub2/api-event/people/domain"
)

type peopleUseCase struct {
	peopleRepository     domain.PeopleRepository
	validationRepository validationsDomain.ValidationRepository
	authRepository       authDomain.AuthRepository
	contextTimeout       time.Duration
	err                  *errDomain.SmartError
}

func NewPeopleUseCase(
	ur domain.PeopleRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	timeout time.Duration,
) domain.PeopleUseCase {
	return &peopleUseCase{
		peopleRepository:     ur,
		validationRepository: validation,
		authRepository:       authRepository,
		contextTimeout:       timeout,
		err:                  errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
