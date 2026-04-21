/*
 * File: `events_usecase.go`
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the EventUseCase interface, which declares methods for interacting with events entities.
 * It includes methods for retrieving, creating, updating, and deleting events data.
 *
 * Last Modified: 2026-04-15
 */

package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type EventUseCase interface {
	GetEvents(ctx context.Context, pagination paramsDomain.PaginationParams, params GetEventsParams) ([]Event,
		*paramsDomain.PaginationResults, error)
	CreateEvent(ctx context.Context, body CreateEventBody) (*string, error)
	UpdateEvent(ctx context.Context, eventId string, body UpdateEventBody) error
	DeleteEvent(ctx context.Context, eventId string) (bool, error)
	GetRolesByEvent(ctx context.Context, eventId string) ([]Role, error)
	EnableDisableEvent(ctx context.Context, eventId string, enable EnableDisableEventRequest) (err error)
}
