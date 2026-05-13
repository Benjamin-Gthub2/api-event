package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

var (
	TypeRegisteredStatus         = "REGISTERED"
	RegistrationRegisteredStatus = "REGISTERED"
	RegistrationPaidStatus       = "PAID"
	RegistrationReceivedStatus   = "RECEIVED"
	RegistrationAttendedStatus   = "ATTENDED"
)

type Registration struct {
	//Description: The id of registration.
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: The flag to verify if qr was sending
	SendQr bool `json:"send_qr" binding:"required" example:"true"`
	//Description: The flag to verify if certificate was sending
	SendCertificate bool `json:"send_certificate" binding:"required" example:"false"`
	//Description: The date of the creation of the registration.
	CreatedAt   *time.Time  `json:"created_at" example:"2023-11-10 08:10:00"`
	Status      Status      `json:"status" binding:"required"`
	Event       Event       `json:"event" binding:"required"`
	Beneficiary Beneficiary `json:"beneficiary" binding:"required"`
	CreatedBy   CreatedBy   `json:"created_by" binding:"required"`
}

type Status struct {
	Id          string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Code        string     `json:"code" binding:"required" example:"REGISTERED"`
	Description string     `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int        `json:"position" example:"1"`
	Enable      bool       `json:"enable" example:"true"`
	CreatedAt   *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
}

type Event struct {
	//Description: the id of event.
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the name of event
	Name string `json:"name" binding:"required" example:"Registration"`
	//Description: the description of event
	Description string `json:"description" binding:"required" example:"Registration Attended"`
	//Description: the date of creation
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type Beneficiary struct {
	//Description: the id of beneficiary
	Id           string       `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	User         User         `json:"user" binding:"required"`
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
	Phone *string `json:"phone" example:"Gthub2"`
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

type User struct {
	//Description: the id of user
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the username
	Username *string  `json:"username" example:"admin.smart"`
	TypeUser TypeUser `json:"type_user" binding:"required"`
}

type TypeUser struct {
	//Description: the id of type user
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the description of type user
	Description *string `json:"description" example:"This is a description"`
	//Description: the code of type user
	Code *string `json:"code" example:"1234"`
	//Description: the date of creation
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreatedBy struct {
	//Description: the id of creator
	Id           string                `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	User         UserCreatedBy         `json:"user" binding:"required"`
	TypeDocument TypeDocumentCreatedBy `json:"type_document" binding:"required"`
	//Description: the document of creator
	Document string `json:"document" binding:"required" example:"73213212"`
	//Description: the names of creator
	Names string `json:"names" binding:"required" example:"Alexander"`
	//Description: the names of creator
	Surname string `json:"surname" binding:"required" example:"Leon"`
	//Description: the last names of creator
	LastName *string `json:"last_name" example:"Gthub2"`
}

type TypeDocumentCreatedBy struct {
	//Description: the id of type document
	Id string `json:"id" binding:"required" example:"1"`
	//Description: the description of type document
	Description string `json:"description" binding:"required" example:"Register Document Identify"`
	//Description: the abbreviated description
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
	//Description: the enable of type document
	Enable string `json:"enable" binding:"required" example:"true"`
}

type UserCreatedBy struct {
	//Description: the id of user
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the username
	Username string            `json:"username" binding:"required" example:"admin.smart"`
	TypeUser TypeUserCreatedBy `json:"type_user" binding:"required"`
}

type TypeUserCreatedBy struct {
	//Description: the id of type user
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the description of type user
	Description string `json:"description" binding:"required" example:"This is a description"`
	//Description: the code of type user
	Code string `json:"code" binding:"required" example:"1234"`
	//Description: the date of creation
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type GetRegistrationsParams struct {
	paramsDomain.Params
	//Description: the initial date
	StartDate *string `json:"start_date" example:"2023-11-10 08:10:00"`
	//Description: the final date
	EndDate *string `json:"end_date" example:"2023-11-10 08:10:00"`
	//Description: the id of the user who created the expense
	CreatedBy *string `json:"created_by" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}

type CreateRegistrationBody struct {
	EventId       string `json:"event_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	BeneficiaryId string `json:"beneficiary_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}

type CreateRegistration struct {
	//Description: the id of registration
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the id of registration
	StatusId string `json:"status_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	//Description: the id of event
	EventId string `json:"event_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the id of beneficiary
	BeneficiaryId string `json:"beneficiary_id" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the id of creator
	CreatedBy string `json:"created_by" binding:"required" example:"200bbbc9-7e93-11ee-89fd-0242ac110016"`
}

type SendQrWhatsAppBody struct {
	//Description: destination phone number with country code, no + sign (e.g. 51987654321)
	PhoneNumber string `json:"phone_number" binding:"required" example:"51987654321"`
}
