package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type MaterialsIssuedUseCase interface {
	GetMaterialIssuedById(ctx context.Context, materialIssuedId string) (*MaterialIssued, error)
	GetMaterialsIssued(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetMaterialsIssuedParams) ([]MaterialIssued, *paramsDomain.PaginationResults, error)
	CreateMaterialIssued(ctx context.Context, userId string, body CreateMaterialIssuedBody) (*string, error)
	UpdateMaterialIssued(ctx context.Context, materialIssuedId string, body UpdateMaterialIssuedBody) error
	DeleteMaterialIssued(ctx context.Context, materialIssuedId string, userId string) error
}
