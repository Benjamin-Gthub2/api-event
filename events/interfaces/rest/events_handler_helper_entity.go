/*
 * File: events_handler_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines helper entities used in the handler layer for managing events related operations.
 *
 * Last Modified: 2026-04-15
 */

package rest

import (
	paginationDomain "github.com/smart0n3/api-shared/params/domain"

	eventsDomain "github.com/Benjamin-Gthub2/api-event/events/domain"
)

type eventsResult struct {
	Data       []eventsDomain.Event               `json:"data" binding:"required"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type deleteEventsResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type rolesResult struct {
	Data   []eventsDomain.Role `json:"data" binding:"required"`
	Status int                 `json:"status" binding:"required"`
}

type eventsSummaryResult struct {
	Data   []eventsDomain.EventSums `json:"data" binding:"required"`
	Status int                      `json:"status" binding:"required"`
}
