package mysql

import "time"

type Session struct {
	Id        string     `db:"session_id"`
	StartDate *time.Time `db:"session_start_date"`
	EndDate   *time.Time `db:"session_end_date"`
	CreatedAt *time.Time `db:"session_created_at"`
	Workshop  Workshop
	CreatedBy CreatedBy
}

type Workshop struct {
	Id        string  `db:"workshop_id"`
	Name      string  `db:"workshop_name"`
	Shortname *string `db:"workshop_shortname"`
	Code      *string `db:"workshop_code"`
	Capacity  int     `db:"workshop_capacity"`
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}

type SessionSums struct {
	Id                 *string    `db:"session_id"`
	StartDate          *time.Time `db:"session_start_date"`
	EndDate            *time.Time `db:"session_end_date"`
	TotalRegistrations *int       `db:"session_total_reg"`
	TotalPayments      *int       `db:"session_total_pay"`
	TotalPresences     *int       `db:"session_total_pres"`
}
