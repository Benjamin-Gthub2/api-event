package mysql

import "time"

type EventType struct {
	Id          string     `db:"event_type_id"`
	Code        string     `db:"event_type_code"`
	Description string     `db:"event_type_description"`
	Enable      bool       `db:"event_type_enable"`
	CreatedAt   *time.Time `db:"event_type_created_at"`
	CreatedBy   CreatedBy
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
