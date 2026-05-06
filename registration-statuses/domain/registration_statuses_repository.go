package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type RegistrationStatusesRepository interface {
	GetRegistrationStatusById(ctx context.Context, registrationStatusId string) (*RegistrationStatus, error)
	GetRegistrationStatuses(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetRegistrationStatusesParams) ([]RegistrationStatus, error)
	GetTotalRegistrationStatuses(ctx context.Context, searchParams GetRegistrationStatusesParams) (*int, error)
	CreateRegistrationStatus(ctx context.Context, body CreateRegistrationStatus) error
	UpdateRegistrationStatus(ctx context.Context, body UpdateRegistrationStatus) error
	DeleteRegistrationStatus(ctx context.Context, body DeleteRegistrationStatus) error
}
