package mysql

import "time"

type Attendance struct {
	Id          string     `db:"attendance_id"`
	CreatedAt   *time.Time `db:"attendance_created_at"`
	Workshop    Workshop
	Beneficiary Beneficiary
	CreatedBy   CreatedBy
}

type Workshop struct {
	Id           string     `db:"workshop_id"`
	Name         string     `db:"workshop_name"`
	Shortname    *string    `db:"workshop_shortname"`
	Code         *string    `db:"workshop_code"`
	Capacity     int        `db:"workshop_capacity"`
	CreatedAt    *time.Time `db:"workshop_created_at"`
	WorkshopType WorkshopType
	Event        Event
}

type WorkshopType struct {
	Id          string `db:"workshop_type_id"`
	Code        string `db:"workshop_type_code"`
	Description string `db:"workshop_type_description"`
}

type Event struct {
	Id   string  `db:"event_id"`
	Name string  `db:"event_name"`
	Code *string `db:"event_code"`
}

type Beneficiary struct {
	Id           string `db:"beneficiary_id"`
	TypeDocument TypeDocument
	Document     string  `db:"beneficiary_document"`
	Names        string  `db:"beneficiary_names"`
	Surname      string  `db:"beneficiary_surname"`
	LastName     *string `db:"beneficiary_last_name"`
}

type TypeDocument struct {
	Id                     string `db:"beneficiary_document_type_id"`
	Description            string `db:"beneficiary_document_type_description"`
	AbbreviatedDescription string `db:"beneficiary_document_type_abbreviated_description"`
	Enable                 string `db:"beneficiary_document_type_enable"`
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
