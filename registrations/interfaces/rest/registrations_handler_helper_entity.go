package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationByIdResult struct {
	Data   *registrationsDomain.Registration `json:"data"`
	Status int                               `json:"status" binding:"required"`
}

type registrationsResult struct {
	Data       []registrationsDomain.Registration `json:"data"`
	Pagination paginationDomain.PaginationResults `json:"pagination" binding:"required"`
	Status     int                                `json:"status" binding:"required"`
}

type registrationsByEventResult struct {
	Data   []registrationsDomain.RegistrationByEvent `json:"data"`
	Status int                                       `json:"status" binding:"required"`
}
