package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type EventTypesRepository interface {
	GetEventTypeById(ctx context.Context, eventTypeId string) (*EventType, error)
	GetEventTypes(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetEventTypesParams) ([]EventType, error)
	GetTotalEventTypes(ctx context.Context, searchParams GetEventTypesParams) (*int, error)
	CreateEventType(ctx context.Context, body CreateEventType) error
	UpdateEventType(ctx context.Context, body UpdateEventType) error
	DeleteEventType(ctx context.Context, body DeleteEventType) error
}
