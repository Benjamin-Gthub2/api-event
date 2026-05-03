package domain

import (
	"time"

	paramsDomain "github.com/smart0n3/api-shared/params/domain"
)

type RegistrationStatus struct {
	Id          string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Code        string     `json:"code" binding:"required" example:"REGISTERED"`
	Description string     `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int        `json:"position" example:"1"`
	Enable      bool       `json:"enable" example:"true"`
	CreatedAt   *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
}

type GetRegistrationStatusesParams struct {
	paramsDomain.Params
}

type CreateRegistrationStatusBody struct {
	Code        string `json:"code" binding:"required" example:"REGISTERED"`
	Description string `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int    `json:"position" binding:"required" example:"1"`
	Enable      bool   `json:"enable" example:"true"`
}

type UpdateRegistrationStatusBody struct {
	Code        string `json:"code" binding:"required" example:"REGISTERED"`
	Description string `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int    `json:"position" binding:"required" example:"1"`
	Enable      bool   `json:"enable" example:"true"`
}

type CreateRegistrationStatus struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	Enable      bool   `json:"enable"`
	CreatedAt   string `json:"created_at"`
}

type UpdateRegistrationStatus struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	Enable      bool   `json:"enable"`
}

type DeleteRegistrationStatus struct {
	Id string `json:"id"`
}
