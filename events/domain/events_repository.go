/*
 * File: events_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the EventRepository interface for events data operations.
 *
 * Last Modified: 2026-14-15
 */

package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type EventRepository interface {
	GetEvents(ctx context.Context, pagination paramsDomain.PaginationParams,
		params GetEventsParams) ([]Event, error)
	GetTotalEvents(ctx context.Context, pagination paramsDomain.PaginationParams,
		params GetEventsParams) (*int, error)
	CreateEvent(ctx context.Context, eventId string, body CreateEventBody) (*string, error)
	UpdateEvent(ctx context.Context, eventId string, body UpdateEventBody) error
	DeleteEvent(ctx context.Context, eventId string) (bool, error)
	GetRolesByEvent(ctx context.Context, eventId string) ([]Role, error)
	EnableDisableEvent(ctx context.Context, eventId string, enable EnableDisableEventRequest) (err error)
	GetEventSums(ctx context.Context, params GetEventSumsParams) ([]EventSums, error)
}
