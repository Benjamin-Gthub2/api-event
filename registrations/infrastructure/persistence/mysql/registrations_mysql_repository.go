/*
 * File: registrations_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewRegistrationsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) registrationsDomain.RegistrationsRepository {
	rep := &registrationsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
