package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	registrationStatusesDomain "github.com/Benjamin-Gthub2/api-event/registration-statuses/domain"
)

type registrationStatusByIdResult struct {
	Data   *registrationStatusesDomain.RegistrationStatus `json:"data"`
	Status int                                            `json:"status" binding:"required"`
}

type registrationStatusesResult struct {
	Data       []registrationStatusesDomain.RegistrationStatus `json:"data"`
	Pagination paginationDomain.PaginationResults              `json:"pagination" binding:"required"`
	Status     int                                             `json:"status" binding:"required"`
}
