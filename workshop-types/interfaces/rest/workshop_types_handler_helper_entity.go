package rest

import (
	paginationDomain "github.com/smart0n3/api-shared/params/domain"

	workshopTypesDomain "github.com/Benjamin-Gthub2/api-event/workshop-types/domain"
)

type workshopTypeByIdResult struct {
	Data   *workshopTypesDomain.WorkshopType `json:"data"`
	Status int                               `json:"status" binding:"required"`
}

type workshopTypesResult struct {
	Data       []workshopTypesDomain.WorkshopType `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}
