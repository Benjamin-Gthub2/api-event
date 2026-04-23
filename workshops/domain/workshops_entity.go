package domain

import (
	"time"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type Workshop struct {
	Id           string       `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name         string       `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname    *string      `json:"shortname" example:"1ER BLOQUE"`
	Code         *string      `json:"code" example:"0001"`
	Capacity     int          `json:"capacity" example:"2"`
	CreatedAt    *time.Time   `json:"created_at" example:"2026-04-21 09:50:04"`
	WorkshopType WorkshopType `json:"workshop_type" binding:"required"`
	Event        Event        `json:"event" binding:"required"`
	CreatedBy    CreatedBy    `json:"created_by" binding:"required"`
}

type WorkshopType struct {
	Id          string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE TALLER"`
}

type Event struct {
	Id   string  `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name string  `json:"name" binding:"required" example:"EVENTO"`
	Code *string `json:"code" example:"0001"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetWorkshopsParams struct {
	paramsDomain.Params
	EventId *string `json:"event_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	TypeId  *string `json:"type_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}

type CreateWorkshopBody struct {
	TypeId    string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
	Code      *string `json:"code" example:"0001"`
	Capacity  int     `json:"capacity" example:"1"`
	EventId   string  `json:"event_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}

type UpdateWorkshopBody struct {
	TypeId    string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
	Code      *string `json:"code" example:"0001"`
	Capacity  int     `json:"capacity" example:"1"`
}

type CreateWorkshop struct {
	Id        string  `json:"id"`
	TypeId    string  `json:"type_id"`
	Name      string  `json:"name"`
	Shortname *string `json:"shortname"`
	Code      *string `json:"code"`
	Capacity  int     `json:"capacity"`
	EventId   string  `json:"event_id"`
	CreatedBy string  `json:"created_by"`
	CreatedAt string  `json:"created_at"`
}

type UpdateWorkshop struct {
	Id        string  `json:"id"`
	TypeId    string  `json:"type_id"`
	Name      string  `json:"name"`
	Shortname *string `json:"shortname"`
	Code      *string `json:"code"`
	Capacity  int     `json:"capacity"`
}

type DeleteWorkshop struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}
