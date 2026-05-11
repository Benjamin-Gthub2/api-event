package rest

import (
	paginationDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"

	workshopSpeakersDomain "github.com/Benjamin-Gthub2/api-event/workshop-speakers/domain"
)

type workshopSpeakerByIdResult struct {
	Data   *workshopSpeakersDomain.WorkshopSpeaker `json:"data"`
	Status int                                     `json:"status" binding:"required"`
}

type workshopSpeakersResult struct {
	Data       []workshopSpeakersDomain.WorkshopSpeaker `json:"data"`
	Pagination paginationDomain.PaginationResults        `json:"pagination" binding:"required"`
	Status     int                                       `json:"status" binding:"required"`
}
