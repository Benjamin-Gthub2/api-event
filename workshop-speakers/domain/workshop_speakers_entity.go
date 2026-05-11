package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type WorkshopSpeaker struct {
	Id                 string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	DegreeAbbreviation *string    `json:"degree_abbreviation" example:"Dr."`
	CreatedAt          *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
	Workshop           Workshop   `json:"workshop" binding:"required"`
	Speaker            Speaker    `json:"speaker" binding:"required"`
	CreatedBy          CreatedBy  `json:"created_by" binding:"required"`
}

type Workshop struct {
	Id        string  `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
}

type Speaker struct {
	Id       string  `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Names    string  `json:"names" binding:"required" example:"Alexander"`
	Surname  string  `json:"surname" binding:"required" example:"Leon"`
	LastName *string `json:"last_name" example:"Gthub2"`
	Document string  `json:"document" binding:"required" example:"73213212"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetWorkshopSpeakersParams struct {
	paramsDomain.Params
	WorkshopId *string `json:"workshop_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	SpeakerId  *string `json:"speaker_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}

type CreateWorkshopSpeakerBody struct {
	WorkshopId         string  `json:"workshop_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	SpeakerId          string  `json:"speaker_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	DegreeAbbreviation *string `json:"degree_abbreviation" example:"Dr."`
}

type CreateWorkshopSpeaker struct {
	Id                 string  `json:"id"`
	WorkshopId         string  `json:"workshop_id"`
	DegreeAbbreviation *string `json:"degree_abbreviation"`
	SpeakerId          string  `json:"speaker_id"`
	CreatedBy          string  `json:"created_by"`
	CreatedAt          string  `json:"created_at"`
}

type DeleteWorkshopSpeaker struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}
