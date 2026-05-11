package rest

import "time"

type createWorkshopValidated struct {
	TypeId    string    `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string    `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string   `json:"shortname" example:"1ER BLOQUE"`
	Code      *string   `json:"code" example:"0001"`
	Capacity  int       `json:"capacity" example:"1"`
	StartDate time.Time `json:"start_date" binding:"required" example:"2026-05-10T18:39:07Z"`
	EndDate   time.Time `json:"end_date" binding:"required" example:"2026-05-10T18:39:13Z"`
	Place     string    `json:"place" binding:"required" example:"QORIKANCHA"`
	EventId   string    `json:"event_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}

type updateWorkshopValidated struct {
	TypeId    string    `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string    `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string   `json:"shortname" example:"1ER BLOQUE"`
	Code      *string   `json:"code" example:"0001"`
	Capacity  int       `json:"capacity" example:"1"`
	StartDate time.Time `json:"start_date" binding:"required" example:"2026-05-10T18:39:07Z"`
	EndDate   time.Time `json:"end_date" binding:"required" example:"2026-05-10T18:39:13Z"`
	Place     string    `json:"place" binding:"required" example:"QORIKANCHA"`
}
