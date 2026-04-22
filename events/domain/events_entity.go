/*
 * File: events_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Defines the EventModel and CreateEventBody structs for events data.
 *
 * Last Modified: 2026-04-15
 */

package domain

import (
	"time"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type Event struct {
	//Description: the id of the event
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the event
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the event
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the code of the event
	Code *string `json:"code" example:"123456789"`
	//Description: the phone of the event
	Phone *string `json:"phone" example:"+1234567890"`
	//Description: the document of the event
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the event
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the event
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//The status enable of the event
	Enable bool `json:"enable" binding:"required" example:"true"`
	//Description: the created_at of the event
	CreatedAt  *time.Time  `json:"created_at" binding:"required" example:"2023-11-10 08:10:00"`
	EventFiles []EventFile `json:"event_files" binding:"required"`
}

type EventFile struct {
	//Description: the id of the event files
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the event files
	Name *string `json:"name" example:"dark"`
	//Description: the description of the event files
	Weight *string `json:"weight" example:"61850"`
	//Description: the phone of the event files
	Url *string `json:"url" example:"event_files/e60372e6-af75-4af2-b3eb-d9598e92dda6.png"`
	//Description: the created_at of the event files
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type CreateEventBody struct {
	//Description: the name of the event
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the event
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the code of the event
	Code string `json:"code" binding:"required" example:"123456789"`
	//Description: the phone of the event
	Phone *string `json:"phone" example:"+1234567890"`
	//Description: the document of the event
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the event
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the event
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//The status enable of the event
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type UpdateEventBody struct {
	//Description: the name of the event
	Name string `json:"name" binding:"required" example:"Odin Corp"`
	//Description: the description of the event
	Description string `json:"description" binding:"required" example:"Proveedor de servicios de mantenimiento"`
	//Description: the code of the event
	Code string `json:"code" binding:"required" example:"123456789"`
	//Description: the phone of the event
	Phone *string `json:"phone" example:"+1234567890"`
	//Description: the document of the event
	Document string `json:"document" binding:"required" example:"123456789"`
	//Description: the address of the event
	Address string `json:"address" binding:"required" example:"123 Main Street"`
	//Description: the industry of the event
	Industry string `json:"industry" binding:"required" example:"Mantenimiento"`
	//The status enable of the event
	Enable bool `json:"enable" binding:"required" example:"true"`
}

type Role struct {
	//Description: the id of the role
	Id string `json:"id" binding:"required" example:"fcdbfacf-8305-11ee-89fd-0242555555"`
	//Description: the name of the role
	Name *string `json:"name" example:"Gerencia"`
	//Description: the description of the role
	Description *string `json:"description" example:"Gerencia del conglomerado"`
	//Description: enable of the role
	Enable *bool `json:"enable" example:"true"`
	//Description: the created_at of the role
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
	Policies  []Policy   `json:"policies" binding:"required"`
	Users     []User     `json:"users" binding:"required"`
}

type Policy struct {
	//Description: the id of the policy
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the name of the policy
	Name string `json:"name" binding:"required" example:"LOGISTICA_REQUERIMIENTOS_CONGLOMERADO"`
	//Description: the description of the policy
	Description string `json:"description" binding:"required" example:"Politica para accesos a logistica requerimientos en todo el conglomerado"`
	//Description: the level of the policy
	Level string `json:"level" binding:"required" example:"system"`
	//Description: enable of the policy
	Enable *bool `json:"enable" example:"true"`
	//Description: the created_at of the policy
	CreatedAt *time.Time     `json:"created_at" example:"2023-11-10 08:10:00"`
	Module    ModuleByPolicy `json:"module"`
}

type ModuleByPolicy struct {
	//Description: the id of the module
	Id *string `json:"id"  example:"739bbbc9-7e93-11ee-89fd-0242ac110018"`
	//Description: the name of the module
	Name *string `json:"name"  example:"Logistic"`
	//Description: the description of the module
	Description *string `json:"description" example:"Modulo de logística"`
	//Description: the code of the module
	Code *string `json:"code"  example:"logistic"`
	//Description: the icon of the module
	Icon *string `json:"icon"  example:"fa fa-home"`
	//Description: The position of the menu user
	Position int `json:"position" binding:"required" example:"1"`
	//Description: The date of created the menu user
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type User struct {
	//Description: the id of the user
	Id         *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	UserRoleId *string `json:"user_role_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	Person     *Person `json:"person"`
}

type Person struct {
	//Description: the id of the person
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	//Description: the document of the person
	Document *string `json:"document" example:"20508997913"`
	//Description: the names of the person
	Names *string `json:"names" example:"Pepo"`
	//Description: the surname of the person
	Surname *string `json:"surname"  example:"Quispe"`
	//Description: the last name of the person
	LastName *string `json:"last_name" example:"Quispe"`
	//Description: the phone of the person
	Phone *string `json:"phone" example:"(01) 317-6000"`
	//Description: the email of the person
	Email *string `json:"email" example:"albert@mail.com"`
	//Description: the gender of the person
	Gender *string `json:"gender" example:"M"`
	//Description: the status of the person
	Enable *bool `json:"enable" example:"true"`
	//Description: the date of created of the person
	CreatedAt    *time.Time    `json:"created_at" example:"2023-11-10 08:10:00"`
	DocumentType *DocumentType `json:"document_type"`
}

type DocumentType struct {
	//Description: the id of the document type
	Id *string `json:"id" example:"25"`
	//Description: the name of the document type
	Number *string `json:"number" example:"25Y"`
	//Description: the abbreviated description of the document type
	Description *string `json:"description" example:"25"`
	//Description: the abbreviated description of the document type
	AbbreviatedDescription *string `json:"abbreviated_description" example:"25"`
}

type EnableDisableEventRequest struct {
	//Description: to enable or disable an event
	Enable *bool `json:"enable" example:"true"`
}

type GetEventsParams struct {
	paramsDomain.Params
	//Description: the status of the event
	Status *bool `json:"status" example:"true"`
	//The name and document number of the event
	NameOrDocument *string `json:"name_or_document" example:"Odin Corp"`
}

type RoleDefault struct {
}
