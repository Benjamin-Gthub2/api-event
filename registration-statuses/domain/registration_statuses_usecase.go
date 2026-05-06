package domain

import (
	"context"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type RegistrationStatusesUseCase interface {
	GetRegistrationStatusById(ctx context.Context, registrationStatusId string) (*RegistrationStatus, error)
	GetRegistrationStatuses(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetRegistrationStatusesParams) ([]RegistrationStatus, *paramsDomain.PaginationResults, error)
	CreateRegistrationStatus(ctx context.Context, body CreateRegistrationStatusBody) (*string, error)
	UpdateRegistrationStatus(ctx context.Context, registrationStatusId string, body UpdateRegistrationStatusBody) error
	DeleteRegistrationStatus(ctx context.Context, registrationStatusId string) error
}
