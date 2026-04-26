package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type WorkshopsUseCase interface {
	GetWorkshopById(ctx context.Context, workshopId string) (*Workshop, error)
	GetWorkshops(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopsParams) ([]Workshop, *paramsDomain.PaginationResults, error)
	CreateWorkshop(ctx context.Context, userId string, body CreateWorkshopBody) (*string, error)
	UpdateWorkshop(ctx context.Context, workshopId string, body UpdateWorkshopBody) error
	DeleteWorkshop(ctx context.Context, workshopId string, userId string) error
	GetWorkshopSummary(ctx context.Context, searchParams GetWorkshopSumsParams) ([]WorkshopSums, error)
}
