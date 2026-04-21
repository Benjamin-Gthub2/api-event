/*
 * File: events_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the use case layer where the layer is initialized.
 *
 * Last Modified: 2026-04-15
 */

package usecase

import (
	"time"

	authDomain "github.com/smart0n3/api-shared/auth/domain"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"
	validationsDomain "github.com/smart0n3/api-shared/validations/domain"

	"github.com/Benjamin-Gthub2/api-event/events/domain"
	rolesDefaultsDomain "github.com/Benjamin-Gthub2/api-event/roles-defaults/domain"
)

type eventsUseCase struct {
	eventsRepository       domain.EventRepository
	validationRepository   validationsDomain.ValidationRepository
	authRepository         authDomain.AuthRepository
	roleDefaultsRepository rolesDefaultsDomain.RoleDefaultsRepository
	contextTimeout         time.Duration
	err                    *errDomain.SmartError
}

func NewEventsUseCase(
	ur domain.EventRepository,
	validation validationsDomain.ValidationRepository,
	authRepository authDomain.AuthRepository,
	roleDefaultsRepository rolesDefaultsDomain.RoleDefaultsRepository,
	timeout time.Duration,
) domain.EventUseCase {
	return &eventsUseCase{
		eventsRepository:       ur,
		validationRepository:   validation,
		authRepository:         authRepository,
		roleDefaultsRepository: roleDefaultsRepository,
		contextTimeout:         timeout,
		err:                    errDomain.NewErr().SetLayer(errDomain.UseCase),
	}
}
