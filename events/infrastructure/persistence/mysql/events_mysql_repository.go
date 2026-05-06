/*
 * File: events_mysql_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file initializes the repository layer to manage event related data.
 *
 * Last Modified: 2026-04-15
 */

package mysql

import (
	"time"

	smartClock "github.com/Benjamin-Gthub2/api-shared/clock"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	eventDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

type eventsMySQLRepo struct {
	clock   smartClock.Clock
	timeout time.Duration
	err     *errDomain.SmartError
}

func NewEventsRepository(
	clock smartClock.Clock,
	mongoTimeout int,
) eventDomain.EventRepository {
	rep := &eventsMySQLRepo{
		clock:   clock,
		timeout: time.Duration(mongoTimeout) * time.Second,
		err:     errDomain.NewErr().SetLayer(errDomain.Infra),
	}
	return rep
}
