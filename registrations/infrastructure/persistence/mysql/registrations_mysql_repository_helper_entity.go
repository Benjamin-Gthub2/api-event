/*
 * File: registrations_mysql_repository_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the entity of the registrations.
 *
 * Last Modified: 2026-04-21
 */

package mysql

import "time"

type Registration struct {
	Id                 string     `db:"registration_id"`
	SendQr             bool       `db:"registration_send_qr"`
	SendCertificate    bool       `db:"registration_send_certificate"`
	CreatedAt          *time.Time `db:"registration_created_at"`
	WorkshopsAttended  int        `db:"workshops_attended"`
	Status             Status
	Event              Event
	Beneficiary        Beneficiary
	CreatedBy          CreatedBy
}

type Status struct {
	Id          string     `db:"status_id"`
	Code        string     `db:"status_code"`
	Description string     `db:"status_description"`
	Position    int        `db:"status_position"`
	Enable      bool       `db:"status_enable"`
	CreatedAt   *time.Time `db:"status_created_at"`
}

type Event struct {
	Id          string     `db:"event_id"`
	Name        string     `db:"event_name"`
	Description string     `db:"event_description"`
	CreatedAt   *time.Time `db:"event_created_at"`
}

type Beneficiary struct {
	Id           string `db:"beneficiary_id"`
	User         User
	TypeDocument TypeDocument
	Document     string  `db:"beneficiary_document"`
	Names        string  `db:"beneficiary_names"`
	Surname      string  `db:"beneficiary_surname"`
	LastName     *string `db:"beneficiary_last_name"`
	Phone        *string `db:"beneficiary_phone"`
}

type TypeDocument struct {
	Id                     string `db:"beneficiary_document_type_id"`
	Description            string `db:"beneficiary_document_type_description"`
	AbbreviatedDescription string `db:"beneficiary_document_type_abbreviated_description"`
	Enable                 string `db:"beneficiary_document_type_enable"`
}

type User struct {
	Id       *string `db:"beneficiary_user_id"`
	Username *string `db:"beneficiary_username"`
	TypeUser TypeUser
}

type TypeUser struct {
	Id          *string    `db:"beneficiary_user_type_id"`
	Description *string    `db:"beneficiary_user_type_description"`
	Code        *string    `db:"beneficiary_user_type_code"`
	CreatedAt   *time.Time `db:"beneficiary_user_type_created_at"`
}

type CreatedBy struct {
	Id           string `db:"creator_id"`
	User         UserCreatedBy
	TypeDocument TypeDocumentCreatedBy
	Document     string  `db:"creator_document"`
	Names        string  `db:"creator_names"`
	Surname      string  `db:"creator_surname"`
	LastName     *string `db:"creator_last_name"`
}

type TypeDocumentCreatedBy struct {
	Id                     string `db:"creator_document_type_id"`
	Description            string `db:"creator_document_type_description"`
	AbbreviatedDescription string `db:"creator_document_type_abbreviated_description"`
	Enable                 string `db:"creator_document_type_enable"`
}

type UserCreatedBy struct {
	Id       string `db:"creator_user_id"`
	Username string `db:"creator_username"`
	TypeUser TypeUserCreatedBy
}

type TypeUserCreatedBy struct {
	Id          string     `db:"creator_user_type_id"`
	Description string     `db:"creator_user_type_description"`
	Code        string     `db:"creator_user_type_code"`
	CreatedAt   *time.Time `db:"creator_user_type_created_at"`
}

type RegistrationByEvent struct {
	Id              string     `db:"registration_by_event_id"`
	SendQr          bool       `db:"registration_by_event_send_qr"`
	SendCertificate bool       `db:"registration_by_event_send_certificate"`
	CreatedAt       *time.Time `db:"registration_by_event_created_at"`
	Status          StatusByEvent
	Event           EventByEvent
	Beneficiary     BeneficiaryByEvent
}

type StatusByEvent struct {
	Id          string     `db:"status_by_event_id"`
	Code        string     `db:"status_by_event_code"`
	Description string     `db:"status_by_event_description"`
	Position    int        `db:"status_by_event_position"`
	Enable      bool       `db:"status_by_event_enable"`
	CreatedAt   *time.Time `db:"status_by_event_created_at"`
}

type EventByEvent struct {
	Id          string     `db:"event_by_event_id"`
	Name        string     `db:"event_by_event_name"`
	Description string     `db:"event_by_event_description"`
	CreatedAt   *time.Time `db:"event_by_event_created_at"`
}

type BeneficiaryByEvent struct {
	Id           string `db:"beneficiary_by_event_id"`
	TypeDocument TypeDocumentByEvent
	Document     string  `db:"beneficiary_by_event_document"`
	Names        string  `db:"beneficiary_by_event_names"`
	Surname      string  `db:"beneficiary_by_event_surname"`
	LastName     *string `db:"beneficiary_by_event_last_name"`
	Phone        *string `db:"beneficiary_by_event_phone"`
}

type TypeDocumentByEvent struct {
	Id                     string `db:"beneficiary_document_type_by_event_id"`
	Description            string `db:"beneficiary_document_type_by_event_description"`
	AbbreviatedDescription string `db:"beneficiary_document_type_abbreviated_by_event_description"`
	Enable                 string `db:"beneficiary_document_type_by_event_enable"`
}
