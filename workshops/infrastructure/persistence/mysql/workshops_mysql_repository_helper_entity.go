package mysql

import "time"

type Workshop struct {
	Id           string     `db:"workshop_id"`
	Name         string     `db:"workshop_name"`
	Shortname    *string    `db:"workshop_shortname"`
	Code         *string    `db:"workshop_code"`
	Capacity     int        `db:"workshop_capacity"`
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
