package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type EventTypesUseCase interface {
	GetEventTypeById(ctx context.Context, eventTypeId string) (*EventType, error)
	GetEventTypes(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetEventTypesParams) ([]EventType, *paramsDomain.PaginationResults, error)
	CreateEventType(ctx context.Context, userId string, body CreateEventTypeBody) (*string, error)
	UpdateEventType(ctx context.Context, eventTypeId string, body UpdateEventTypeBody) error
	DeleteEventType(ctx context.Context, eventTypeId string, userId string) error
}
