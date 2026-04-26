package rest

import (
	paginationDomain "github.com/smart0n3/api-shared/params/domain"

	sessionsDomain "github.com/Benjamin-Gthub2/api-event/sessions/domain"
)

type sessionByIdResult struct {
	Data   *sessionsDomain.Session `json:"data"`
	Status int                     `json:"status" binding:"required"`
}

type sessionsResult struct {
	Data       []sessionsDomain.Session           `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type sessionsSummaryResult struct {
	Data   []sessionsDomain.SessionSums `json:"data" binding:"required"`
	Status int                          `json:"status" binding:"required"`
}
