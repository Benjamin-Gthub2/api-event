package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type WorkshopSpeakersRepository interface {
	GetWorkshopSpeakerById(ctx context.Context, workshopSpeakerId string) (*WorkshopSpeaker, error)
	GetWorkshopSpeakers(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopSpeakersParams) ([]WorkshopSpeaker, error)
	GetTotalWorkshopSpeakers(ctx context.Context, searchParams GetWorkshopSpeakersParams) (*int, error)
	CreateWorkshopSpeaker(ctx context.Context, body CreateWorkshopSpeaker) error
	DeleteWorkshopSpeaker(ctx context.Context, body DeleteWorkshopSpeaker) error
}
