package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	eventTypesDomain "github.com/Benjamin-Gthub2/api-event/event-types/domain"
)

type eventTypeByIdResult struct {
	Data   *eventTypesDomain.EventType `json:"data"`
	Status int                         `json:"status" binding:"required"`
}

type eventTypesResult struct {
	Data       []eventTypesDomain.EventType       `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
