package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type WorkshopSpeakersUseCase interface {
	GetWorkshopSpeakerById(ctx context.Context, workshopSpeakerId string) (*WorkshopSpeaker, error)
	GetWorkshopSpeakers(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopSpeakersParams) ([]WorkshopSpeaker, *paramsDomain.PaginationResults, error)
	CreateWorkshopSpeaker(ctx context.Context, userId string, body CreateWorkshopSpeakerBody) (*string, error)
	DeleteWorkshopSpeaker(ctx context.Context, workshopSpeakerId string, userId string) error
}
