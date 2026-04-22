/*
 * File: people_handler_helper_validation.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package rest

type createPersonValidate struct {
	UserId         *string `json:"user_id" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	TypeDocumentId string  `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	Document       string  `json:"document" binding:"required" example:"77895428"`
	Names          string  `json:"names" binding:"required" example:"LUCY ANDREA"`
	Surname        string  `json:"surname" binding:"required" example:"HANCCO"`
	LastName       *string `json:"last_name" example:"HUILLCA"`
	Phone          *string `json:"phone" example:"918547496"`
	Email          *string `json:"email" example:"lucyhancco@gmail.com"`
	Gender         *string `json:"gender" example:"MASCULINO"`
	Enable         bool    `json:"enable" example:"1"`
}

type UpdatePersonValidate struct {
	UserId         *string `json:"user_id"`
	TypeDocumentId string  `json:"type_document_id" binding:"required" example:"00a58522-93b4-11ee-a040-0242ac11000e"`
	Document       string  `json:"document" binding:"required" example:"77895428"`
	Names          string  `json:"names" binding:"required" example:"LUCY ANDREA"`
	Surname        string  `json:"surname" binding:"required" example:"HANCCO"`
	LastName       *string `json:"last_name" example:"HUILLCA"`
	Phone          *string `json:"phone" example:"918547496"`
	Email          *string `json:"email" example:"lucyhancco@gmail.com"`
	Gender         *string `json:"gender" example:"MASCULINO"`
	Enable         bool    `json:"enable" example:"1"`
}
