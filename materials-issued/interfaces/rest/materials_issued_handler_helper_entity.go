package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	materialsIssuedDomain "github.com/Benjamin-Gthub2/api-event/materials-issued/domain"
)

type materialIssuedByIdResult struct {
	Data   *materialsIssuedDomain.MaterialIssued `json:"data"`
	Status int                                   `json:"status" binding:"required"`
}

type materialsIssuedResult struct {
	Data       []materialsIssuedDomain.MaterialIssued `json:"data"`
	Pagination paginationDomain.PaginationResults     `json:"pagination" binding:"required"`
	Status     int                                    `json:"status" binding:"required"`
}
