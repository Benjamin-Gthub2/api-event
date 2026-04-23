package mysql

import "time"

type WorkshopType struct {
	Id          string     `db:"workshop_type_id"`
	Code        string     `db:"workshop_type_code"`
	Description string     `db:"workshop_type_description"`
	Enable      bool       `db:"workshop_type_enable"`
	CreatedAt   *time.Time `db:"workshop_type_created_at"`
	CreatedBy   CreatedBy
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
