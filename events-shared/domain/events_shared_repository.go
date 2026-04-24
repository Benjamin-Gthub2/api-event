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
	UpdateEventTotals(ctx context.Context, tx *sql.Tx, paymentId string, body UpdateEventTotals) error
	UpdateWorkshopTotals(ctx context.Context, tx *sql.Tx, paymentId string, body UpdateWorkshopTotals) error
	UpdateSessionTotals(ctx context.Context, tx *sql.Tx, paymentId string, body UpdateSessionTotals) error
}
