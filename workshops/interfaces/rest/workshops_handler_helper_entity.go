package rest

import (
	paginationDomain "github.com/smart0n3/api-shared/params/domain"

	workshopsDomain "github.com/Benjamin-Gthub2/api-event/workshops/domain"
)

type workshopByIdResult struct {
	Data   *workshopsDomain.Workshop `json:"data"`
	Status int                       `json:"status" binding:"required"`
}

type workshopsResult struct {
	Data       []workshopsDomain.Workshop         `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type workshopsSummaryResult struct {
	Data   []workshopsDomain.WorkshopSums `json:"data" binding:"required"`
	Status int                            `json:"status" binding:"required"`
}
