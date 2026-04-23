/*
 * File: events_mysql_repository_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the eventModel entity.
 *
 * Last Modified: 2026-04-15
 */

package mysql

import "time"

type EventHelper struct {
	Id          string     `db:"events_id" `
	Name        string     `db:"events_name"`
	Description string     `db:"events_description"`
	Code        *string    `db:"events_code"`
	Phone       string     `db:"events_phone"`
	Document    string     `db:"events_document"`
	Address     string     `db:"events_address"`
	Industry    string     `db:"events_industry"`
	Enable      bool       `db:"events_enable"`
	CreatedAt   *time.Time `db:"events_created_at"`
	EventFiles  []EventFileHelper
}

type EventFileHelper struct {
	Id        *string    `db:"events_files_id"`
	Name      *string    `db:"events_files_name"`
	Weight    *string    `db:"events_files_weight"`
	Url       *string    `db:"events_files_url"`
	CreatedAt *time.Time `db:"events_files_created_at"`
}

type RoleByEvent struct {
	Id          string     `db:"role_id"`
	Name        *string    `db:"role_name"`
	Description *string    `db:"role_description"`
	Enable      *bool      `db:"role_enable"`
	CreatedAt   *time.Time `db:"role_created_at"`
	Policies    []Policy
	Users       []User
}

type Policy struct {
	Id          string     `db:"policy_id"`
	Name        string     `db:"policy_name"`
	Description string     `db:"policy_description"`
	Level       string     `db:"policy_level"`
	Enable      *bool      `db:"policy_enable"`
	CreatedAt   *time.Time `db:"policy_created_at"`
	Module      ModuleByPolicy
}

type ModuleByPolicy struct {
	Id          *string    `db:"module_id"`
	Name        *string    `db:"module_name"`
	Description *string    `db:"module_description"`
	Code        *string    `db:"module_code"`
	Icon        *string    `db:"module_icon"`
	Position    int        `db:"module_position"`
	CreatedAt   *time.Time `db:"module_created_at"`
}

type User struct {
	Id         *string `db:"user_id"`
	UserRoleId *string `db:"user_role_id" example:"739bbbc9-7e93-11ee-89fd-0242ac113421"`
	Person     *Person
}

type Person struct {
	Id           *string    `db:"person_id"`
	Document     *string    `db:"person_document"`
	Names        *string    `db:"person_names"`
	Surname      *string    `db:"person_surname"`
	LastName     *string    `db:"person_last_name"`
	Phone        *string    `db:"person_phone"`
	Email        *string    `db:"person_email"`
	Gender       *string    `db:"person_gender"`
	Enable       *bool      `db:"person_enable"`
	CreatedAt    *time.Time `db:"person_created_at"`
	DocumentType *DocumentType
}

type DocumentType struct {
	Id                     *string `db:"document_type_id"`
	Number                 *string `db:"document_type_number"`
	Description            *string `db:"document_type_description"`
	AbbreviatedDescription *string `db:"document_type_abbreviated_description"`
}

type EventSums struct {
	Id                 string `db:"event_id"`
	Name               string `db:"event_name"`
	TotalRegistrations int    `db:"event_total_reg"`
	TotalPayments      int    `db:"event_total_pay"`
	TotalPresences     int    `db:"event_total_pres"`
	WorkshopSums       []WorkshopSums
}

type WorkshopSums struct {
	Id                 *string `db:"workshop_id"`
	Name               *string `db:"workshop_name"`
	TotalRegistrations *int    `db:"workshop_total_reg"`
	TotalPayments      *int    `db:"workshop_total_pay"`
	TotalPresences     *int    `db:"workshop_total_pres"`
	SessionSums        []SessionSums
}

type SessionSums struct {
	Id                 *string    `db:"session_id"`
	StartDate          *time.Time `db:"session_start_date"`
	EndDate            *time.Time `db:"session_end_date"`
	TotalRegistrations *int       `db:"session_total_reg"`
	TotalPayments      *int       `db:"session_total_pay"`
	TotalPresences     *int       `db:"session_total_pres"`
}
