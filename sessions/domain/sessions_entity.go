package domain

import (
	"time"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type Session struct {
	Id        string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate *time.Time `json:"start_date" example:"2026-04-21 09:51:23"`
	EndDate   *time.Time `json:"end_date" example:"2026-04-21 11:00:00"`
	CreatedAt *time.Time `json:"created_at" example:"2026-04-21 09:51:57"`
	Workshop  Workshop   `json:"workshop" binding:"required"`
	CreatedBy CreatedBy  `json:"created_by" binding:"required"`
}

type Workshop struct {
	Id        string  `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
	Code      *string `json:"code" example:"0001"`
	Capacity  int     `json:"capacity" example:"1"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetSessionsParams struct {
	paramsDomain.Params
	WorkshopId *string `json:"workshop_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate  *string `json:"start_date" example:"2026-04-21"`
	EndDate    *string `json:"end_date" example:"2026-04-21"`
}

type CreateSessionBody struct {
	WorkshopId string `json:"workshop_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate  string `json:"start_date" binding:"required" example:"2026-04-21 09:00:00"`
	EndDate    string `json:"end_date" binding:"required" example:"2026-04-21 11:00:00"`
}

type UpdateSessionBody struct {
	StartDate string `json:"start_date" binding:"required" example:"2026-04-21 09:00:00"`
	EndDate   string `json:"end_date" binding:"required" example:"2026-04-21 11:00:00"`
}

type CreateSession struct {
	Id         string `json:"id"`
	WorkshopId string `json:"workshop_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
}

type UpdateSession struct {
	Id        string `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type DeleteSession struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}

type GetSessionSumsParams struct {
	paramsDomain.Params
	//Description: the id of session
	SessionId *string `json:"session_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
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
