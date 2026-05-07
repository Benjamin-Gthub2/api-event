package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type MaterialsIssuedRepository interface {
	GetMaterialIssuedById(ctx context.Context, materialIssuedId string) (*MaterialIssued, error)
	GetMaterialsIssued(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetMaterialsIssuedParams) ([]MaterialIssued, error)
	GetTotalMaterialsIssued(ctx context.Context, searchParams GetMaterialsIssuedParams) (*int, error)
	CreateMaterialIssued(ctx context.Context, body CreateMaterialIssued) error
	UpdateMaterialIssued(ctx context.Context, body UpdateMaterialIssued) error
	DeleteMaterialIssued(ctx context.Context, body DeleteMaterialIssued) error
}
