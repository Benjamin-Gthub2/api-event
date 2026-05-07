package mysql

import "time"

type RegistrationStatus struct {
	Id          string     `db:"registration_status_id"`
	Code        string     `db:"registration_status_code"`
	Description string     `db:"registration_status_description"`
	Position    int        `db:"registration_status_position"`
	Enable      bool       `db:"registration_status_enable"`
	CreatedAt   *time.Time `db:"registration_status_created_at"`
}
