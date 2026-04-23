package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type SessionsRepository interface {
	GetSessionById(ctx context.Context, sessionId string) (*Session, error)
	GetSessions(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetSessionsParams) ([]Session, error)
	GetTotalSessions(ctx context.Context, searchParams GetSessionsParams) (*int, error)
	CreateSession(ctx context.Context, body CreateSession) error
	UpdateSession(ctx context.Context, body UpdateSession) error
	DeleteSession(ctx context.Context, body DeleteSession) error
}
