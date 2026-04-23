package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type WorkshopTypesRepository interface {
	GetWorkshopTypeById(ctx context.Context, workshopTypeId string) (*WorkshopType, error)
	GetWorkshopTypes(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopTypesParams) ([]WorkshopType, error)
	GetTotalWorkshopTypes(ctx context.Context, searchParams GetWorkshopTypesParams) (*int, error)
	CreateWorkshopType(ctx context.Context, body CreateWorkshopType) error
	UpdateWorkshopType(ctx context.Context, body UpdateWorkshopType) error
	DeleteWorkshopType(ctx context.Context, body DeleteWorkshopType) error
}
