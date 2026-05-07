package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type WorkshopType struct {
	Id          string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Code        string     `json:"code" binding:"required" example:"0001"`
	Description string     `json:"description" binding:"required" example:"TIPO DE TALLER"`
	Enable      bool       `json:"enable" example:"true"`
	CreatedAt   *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
	CreatedBy   CreatedBy  `json:"created_by" binding:"required"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetWorkshopTypesParams struct {
	paramsDomain.Params
}

type CreateWorkshopTypeBody struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE TALLER"`
	Enable      bool   `json:"enable" example:"true"`
}

type UpdateWorkshopTypeBody struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE TALLER"`
	Enable      bool   `json:"enable" example:"true"`
}

type CreateWorkshopType struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Enable      bool   `json:"enable"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
}

type UpdateWorkshopType struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Enable      bool   `json:"enable"`
}

type DeleteWorkshopType struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
	DeletedAt string `json:"deleted_at"`
}
