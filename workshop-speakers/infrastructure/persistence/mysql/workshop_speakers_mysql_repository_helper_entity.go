package mysql

import "time"

type WorkshopSpeaker struct {
	Id                 string     `db:"workshop_speaker_id"`
	DegreeAbbreviation *string    `db:"workshop_speaker_degree_abbreviation"`
	CreatedAt          *time.Time `db:"workshop_speaker_created_at"`
	Workshop           Workshop
	Speaker            Speaker
	CreatedBy          CreatedBy
}

type Workshop struct {
	Id        string  `db:"workshop_id"`
	Name      string  `db:"workshop_name"`
	Shortname *string `db:"workshop_shortname"`
}

type Speaker struct {
	Id       string  `db:"speaker_id"`
	Names    string  `db:"speaker_names"`
	Surname  string  `db:"speaker_surname"`
	LastName *string `db:"speaker_last_name"`
	Document string  `db:"speaker_document"`
}

type CreatedBy struct {
	Id       string `db:"created_by_id"`
	Username string `db:"created_by_username"`
}
