package domain

import (
	"context"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type SessionsUseCase interface {
	GetSessionById(ctx context.Context, sessionId string) (*Session, error)
	GetSessions(ctx context.Context, pagination paramsDomain.PaginationParams, searchParams GetSessionsParams) ([]Session, *paramsDomain.PaginationResults, error)
	CreateSession(ctx context.Context, userId string, body CreateSessionBody) (*string, error)
	UpdateSession(ctx context.Context, sessionId string, body UpdateSessionBody) error
	DeleteSession(ctx context.Context, sessionId string, userId string) error
}
