package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type Workshop struct {
	Id           string       `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name         string       `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname    *string      `json:"shortname" example:"1ER BLOQUE"`
	Code         *string      `json:"code" example:"0001"`
	Capacity     int          `json:"capacity" example:"2"`
	StartDate    *time.Time   `json:"start_date" example:"2026-05-10T18:39:07Z"`
	EndDate      *time.Time   `json:"end_date" example:"2026-05-10T18:39:13Z"`
	Place        string       `json:"place" example:"QORIKANCHA"`
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
	EventId     *string `json:"event_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	TypeId      *string `json:"type_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate   *string `json:"start_date" example:"2026-05-15"`
	OnlyToday   *bool   `json:"only_today" example:"true"`
	SearchValue *string `json:"searchvalue" example:"TALLER"`
}

type CreateWorkshopBody struct {
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

type UpdateWorkshopBody struct {
	TypeId    string    `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string    `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string   `json:"shortname" example:"1ER BLOQUE"`
	Code      *string   `json:"code" example:"0001"`
	Capacity  int       `json:"capacity" example:"1"`
	StartDate time.Time `json:"start_date" binding:"required" example:"2026-05-10T18:39:07Z"`
	EndDate   time.Time `json:"end_date" binding:"required" example:"2026-05-10T18:39:13Z"`
	Place     string    `json:"place" binding:"required" example:"QORIKANCHA"`
}

type CreateWorkshop struct {
	Id        string  `json:"id"`
	TypeId    string  `json:"type_id"`
	Name      string  `json:"name"`
	Shortname *string `json:"shortname"`
	Code      *string `json:"code"`
	Capacity  int     `json:"capacity"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Place     string  `json:"place"`
	EventId   string  `json:"event_id"`
	CreatedBy string  `json:"created_by"`
	CreatedAt string  `json:"created_at"`
}

type UpdateWorkshop struct {
	Id        string    `json:"id"`
	TypeId    string    `json:"type_id"`
	Name      string    `json:"name"`
	Shortname *string   `json:"shortname"`
	Code      *string   `json:"code"`
	Capacity  int       `json:"capacity"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Place     string    `json:"place"`
}

type DeleteWorkshop struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}

type GetWorkshopSumsParams struct {
	paramsDomain.Params
	//Description: the id of workshop
	WorkshopId *string `json:"workshop_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	//Description: the value of workshop
	SearchValue *string `json:"search_value" example:"TALLER"`
	//Description: the start date
	StartDate *string `json:"start_date" example:"2026-04-21"`
	//Description: the end date
	EndDate *string `json:"end_date" example:"2026-04-21"`
}

type WorkshopSums struct {
	//Description: the id of workshop
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the name of workshop
	Name *string `json:"name" example:"Workshop 1"`
	//Description: the name of workshop
	StartDate *time.Time `json:"start_date" example:"2023-11-10 08:10:00"`
	//Description: the name of workshop
	EndDate *time.Time `json:"end_date" example:"2023-11-10 08:10:00"`
	//Description: the place of workshop
	Place *string `json:"place" example:"Qorikancha"`
	//Description: the capacity of workshop
	Capacity *int `json:"capacity" example:"1"`
	//Description: the quantity of presences
	TotalPresences *int `json:"total_presences" example:"1"`
}

type Speaker struct {
	//Description: the degree abbreviation of speaker
	DegreeAbbreviation *string `json:"degree_abbreviation" example:"Dr."`
	//Description: the id of speaker
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the name of speaker
	Name *string `json:"name" example:"Pepe"`
	//Description: the surname of speaker
	Surname *string `json:"surname" example:"Quintana"`
	//Description: the lastname of speaker
	LastName *string `json:"last_name" example:"Garcia"`
}

type SessionSums struct {
	//Description: the id of session
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the name of session
	StartDate *time.Time `json:"start_date" example:"2023-11-10 08:10:00"`
	//Description: the name of session
	EndDate *time.Time `json:"end_date" example:"2023-11-10 08:10:00"`
	//Description: the quantity of registrations
	TotalRegistrations *int `json:"total_registrations" example:"1"`
	//Description: the quantity of payments
	TotalPayments *int `json:"total_payments" example:"1"`
	//Description: the quantity of presences
	TotalPresences *int `json:"total_presences" example:"1"`
}
