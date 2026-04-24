/*
 * File: events_shared_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository.
 *
 * Last Modified: 2026-04-24
 */

package mysql

import (
	"time"

	smartClock "github.com/smart0n3/api-shared/clock"
	errDomain "github.com/smart0n3/api-shared/error-core/domain"

	eventSharedDomain "github.com/Benjamin-Gthub2/api-event/events-shared/domain"
)

type eventSharedMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewEventSharedRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) eventSharedDomain.EventSharedRepository {
	rep := &eventSharedMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
