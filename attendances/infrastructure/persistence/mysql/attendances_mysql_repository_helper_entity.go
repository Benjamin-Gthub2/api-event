package mysql

import "time"

type Attendance struct {
	Id        string     `db:"attendance_id"`
	CreatedAt *time.Time `db:"attendance_created_at"`
	CreatedBy CreatedBy
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
