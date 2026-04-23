package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type WorkshopTypesUseCase interface {
	GetWorkshopTypeById(ctx context.Context, workshopTypeId string) (*WorkshopType, error)
	GetWorkshopTypes(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetWorkshopTypesParams) ([]WorkshopType, *paramsDomain.PaginationResults, error)
	CreateWorkshopType(ctx context.Context, userId string, body CreateWorkshopTypeBody) (*string, error)
	UpdateWorkshopType(ctx context.Context, workshopTypeId string, body UpdateWorkshopTypeBody) error
	DeleteWorkshopType(ctx context.Context, workshopTypeId string, userId string) error
}
