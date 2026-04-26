package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type WorkshopsRepository interface {
	GetWorkshopById(ctx context.Context, workshopId string) (*Workshop, error)
	GetWorkshops(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopsParams) ([]Workshop, error)
	GetTotalWorkshops(ctx context.Context, searchParams GetWorkshopsParams) (*int, error)
	CreateWorkshop(ctx context.Context, body CreateWorkshop) error
	UpdateWorkshop(ctx context.Context, body UpdateWorkshop) error
	DeleteWorkshop(ctx context.Context, body DeleteWorkshop) error
	GetWorkshopSums(ctx context.Context, searchParams GetWorkshopSumsParams) ([]WorkshopSums, error)
}
