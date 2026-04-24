/*
 * File: events_shared_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository interface in the domain.
 *
 * Last Modified: 2026-04-24
 */

package domain

import (
	"context"
	"database/sql"
)

type EventSharedRepository interface {
	UpdateEventTotals(ctx context.Context, tx *sql.Tx, eventId string, body UpdateEventTotals) error
	UpdateWorkshopTotals(ctx context.Context, tx *sql.Tx, workshopId string, body UpdateWorkshopTotals) error
	UpdateSessionTotals(ctx context.Context, tx *sql.Tx, sessionId string, body UpdateSessionTotals) error
	GetEventTotals(ctx context.Context, tx *sql.Tx, eventId string) (*EventTotals, error)
	GetWorkshopTotals(ctx context.Context, tx *sql.Tx, workshopId string) (*WorkshopTotals, error)
	GetSessionTotals(ctx context.Context, tx *sql.Tx, sessionId string) (*SessionTotals, error)
}
