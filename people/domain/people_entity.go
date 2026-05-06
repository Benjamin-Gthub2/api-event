/*
 * File: people_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type People struct {
	//Description: the id of the people
	Id string `json:"id" binding:"required" example:"0abbb86f-9836-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document string `json:"document" binding:"required" example:"77895428"`
	//Description: the name of the people
	Names string `json:"names" binding:"required" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname string `json:"surname" binding:"required" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" binding:"required" example:"HUILLCA"`
	//Description: the phone of the people
	Phone *string `json:"phone" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email" binding:"required" example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" binding:"required" example:"MASCULINO"`
	//Description: the status of the people
	Enable bool `json:"enable" binding:"required" example:"1"`
	//Description: the date of created of the people
	CreatedAt    *time.Time   `json:"created_at" example:"2023-11-10 08:10:00"`
	User         *User        `json:"user"`
	DocumentType DocumentType `json:"document_type" binding:"required"`
}

type DocumentType struct {
	//Description: the id of the document type
	Id string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the status of the document type
	Description string `json:"description" binding:"required" example:"DOCUMENTO NACIONAL DE IDENTIDAD"`
	//Description: the abbreviated description of the document type
	AbbreviatedDescription string `json:"abbreviated_description" binding:"required" example:"DNI"`
}

type User struct {
	//Description: the id of the user
	Id *string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110016"`
	//Description: the username of the user
	UserName *string `json:"username" example:"pepito.quispe@smartc.pe"`
	//Description: the date of created of the user
	CreatedAt *time.Time `json:"created_at" example:"2023-11-10 08:10:00"`
}

type GetPeopleParams struct {
	paramsDomain.Params
	//Description: search name of the people
	SearchName *string `json:"search_name" example:"LUCY ANDREA"`
	//Description: the document number id of the people
	Document *string `json:"document" example:"77895428"`
	//Description: the abbreviated description of the document type
	DocumentTypeId string `json:"document_type_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
}

type CreatePersonBody struct {
	//Description: the id of the user
	UserId *string `json:"user_id"`
	//Description: the type of the document
	TypeDocumentId string `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document string `json:"document" binding:"required" example:"77895428"`
	//Description: the name of the people
	Names string `json:"names" binding:"required" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname string `json:"surname" binding:"required" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" example:"HUILLCA"`
	//Description: the phone of the people
	Phone *string `json:"phone" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email"  example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" example:"MASCULINO"`
	//Description: the status of the people
	Enable bool `json:"enable" binding:"required" example:"1"`
}

type UpdatePersonBody struct {
	//Description: the id of the user
	UserId *string `json:"user_id"`
	//Description: the type of the document
	TypeDocumentId string `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	//Description: the document number of the people
	Document string `json:"document" binding:"required" example:"77895428"`
	//Description: the name of the people
	Names string `json:"names" binding:"required" example:"LUCY ANDREA"`
	//Description: the surname of the people
	Surname string `json:"surname" binding:"required" example:"HANCCO"`
	//Description: the last name of the people
	LastName *string `json:"last_name" example:"HUILLCA"`
	//Description: the phone of the people
	Phone *string `json:"phone" example:"918547496"`
	//Description: the email of the people
	Email *string `json:"email" example:"lucyhancco@gmail.com"`
	//Description: the gender of the people
	Gender *string `json:"gender" example:"MASCULINO"`
	//Description: the status of the people
	Enable bool `json:"enable" binding:"required" example:"1"`
}
