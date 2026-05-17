package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type Attendance struct {
	Id          string      `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	CreatedAt   *time.Time  `json:"created_at" example:"2026-04-21 09:50:04"`
	Workshop    Workshop    `json:"workshop" binding:"required"`
	Beneficiary Beneficiary `json:"beneficiary" binding:"required"`
	CreatedBy   CreatedBy   `json:"created_by" binding:"required"`
}

type Workshop struct {
	Id           string       `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name         string       `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname    *string      `json:"shortname" example:"1ER BLOQUE"`
	Code         *string      `json:"code" example:"0001"`
	Capacity     int          `json:"capacity" example:"2"`
	CreatedAt    *time.Time   `json:"created_at" example:"2026-04-21 09:50:04"`
	WorkshopType WorkshopType `json:"workshop_type" binding:"required"`
	Event        Event        `json:"event" binding:"required"`
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

type Beneficiary struct {
	//Description: the id of beneficiary
	Id           string       `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	TypeDocument TypeDocument `json:"type_document" binding:"required"`
	//Description: the document of beneficiary
	Document string `json:"document" binding:"required" example:"73213212"`
	//Description: the names of beneficiary
	Names string `json:"names" binding:"required" example:"Alexander"`
	//Description: the names of beneficiary
	Surname string `json:"surname" binding:"required" example:"Leon"`
	//Description: the last names of beneficiary
	LastName *string `json:"last_name" example:"Gthub2"`
	//Description: the phone of beneficiary
	Phone *string `json:"phone" example:"73213212"`
}

type TypeDocument struct {
	//Description: the id of type document
	Id string `json:"id" binding:"required" example:"1"`
	//Description: the description of type document
	Description string `json:"description" binding:"required" example:"Register Document Identify"`
	//Description: the abbreviated description
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	//Description: the enable of type document
	Enable string `json:"enable" binding:"required" example:"true"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetAttendancesParams struct {
	paramsDomain.Params
	EventId       *string `json:"event_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	WorkshopId    *string `json:"workshop_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	BeneficiaryId *string `json:"beneficiary_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate     *string `json:"start_date" example:"2026-01-01"`
	EndDate       *string `json:"end_date" example:"2026-12-31"`
	SearchValue   *string `json:"searchvalue" example:"Alexander"`
}

type CreateAttendanceBody struct {
	WorkshopId    string `json:"workshop_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	BeneficiaryId string `json:"beneficiary_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}

type CreateAttendance struct {
	Id            string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	WorkshopId    string `json:"workshop_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	BeneficiaryId string `json:"beneficiary_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	CreatedBy     string `json:"created_by" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}

type DeleteAttendance struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
}
