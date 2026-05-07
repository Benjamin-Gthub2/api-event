/*
 * File: people_mysql_repository_helper_entity.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-12
 */

package mysql

import (
	"time"
)

type People struct {
	Id           string     `db:"people_id"`
	Document     string     `db:"people_document"`
	Names        string     `db:"people_name"`
	Surname      string     `db:"people_surname"`
	LastName     *string    `db:"people_last_name"`
	Phone        *string    `db:"people_phone"`
	Email        *string    `db:"people_email"`
	Gender       *string    `db:"people_gender"`
	Enable       bool       `db:"people_enable"`
	CreatedAt    *time.Time `db:"people_created_at"`
	User         *User
	DocumentType DocumentType
}

type User struct {
	Id        *string    `db:"user_id"`
	UserName  *string    `db:"user_username"`
	CreatedAt *time.Time `db:"user_created_at"`
}

type DocumentType struct {
	Id                     string `db:"document_type_id"`
	Description            string `db:"document_type_description"`
	AbbreviatedDescription string `db:"document_type_abbreviated_description"`
}
