package mysql

import "time"

type Workshop struct {
	Id           string     `db:"workshop_id"`
	Name         string     `db:"workshop_name"`
	Shortname    *string    `db:"workshop_shortname"`
	Code         *string    `db:"workshop_code"`
	Capacity     int        `db:"workshop_capacity"`
	StartDate    *time.Time `db:"workshop_start_date"`
	EndDate      *time.Time `db:"workshop_end_date"`
	Place        string     `db:"workshop_place"`
	CreatedAt    *time.Time `db:"workshop_created_at"`
	WorkshopType WorkshopType
	Event        Event
	CreatedBy    CreatedBy
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

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}

type WorkshopSums struct {
	Id             *string    `db:"workshop_id"`
	Name           *string    `db:"workshop_name"`
	StartDate      *time.Time `db:"workshop_start_date"`
	EndDate        *time.Time `db:"workshop_end_date"`
	Place          *string    `db:"workshop_place"`
	Capacity       *int       `db:"workshop_capacity"`
	TotalPresences *int       `db:"workshop_total_pres"`
	Speaker        []Speaker
}

type Speaker struct {
	DegreeAbbreviation *string `db:"workshop_speaker_degree_abbreviation"`
	Id                 *string `db:"speaker_id"`
	Name               *string `db:"speaker_name"`
	Surname            *string `db:"speaker_surname"`
	LastName           *string `db:"speaker_last_name"`
}

type SessionSums struct {
	Id                 *string    `db:"session_id"`
	StartDate          *time.Time `db:"session_start_date"`
	EndDate            *time.Time `db:"session_end_date"`
	TotalRegistrations *int       `db:"session_total_reg"`
	TotalPayments      *int       `db:"session_total_pay"`
	TotalPresences     *int       `db:"session_total_pres"`
}
