package domain

import (
	"time"
)

type RegistrationStatus struct {
	Id          string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Code        string     `json:"code" binding:"required" example:"REGISTERED"`
	Description string     `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int        `json:"position" example:"1"`
	Enable      bool       `json:"enable" example:"true"`
	CreatedAt   *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
}
