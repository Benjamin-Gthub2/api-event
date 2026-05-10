/*
 * File: attendances_rt_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the repository of the attendances.
 *
 * Last Modified: 2026-04-22
 */

package domain

import (
	"context"
)

type AttendancesRTRepository interface {
	SendNotification(ctx context.Context, topic string, message string) error
}
