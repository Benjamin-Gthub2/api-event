package mysql

import "time"

type MaterialIssued struct {
	Id          string     `db:"material_issued_id"`
	Description *string    `db:"material_issued_description"`
	CreatedAt   *time.Time `db:"material_issued_created_at"`
	CreatedBy   CreatedBy
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
