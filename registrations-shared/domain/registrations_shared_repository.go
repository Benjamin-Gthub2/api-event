package domain

import (
	"context"
)

type RegistrationSharedRepository interface {
	GetStatusByCode(ctx context.Context, code string) (*RegistrationStatus, error)
}
