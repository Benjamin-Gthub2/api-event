package domain

import (
	"time"

	paramsDomain "github.com/Benjamin-Gthub2/api-shared/params/domain"
)

type MaterialIssued struct {
	Id          string     `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Description *string    `json:"description" example:"Kit de materiales"`
	CreatedAt   *time.Time `json:"created_at" example:"2026-04-21 09:50:04"`
	CreatedBy   CreatedBy  `json:"created_by" binding:"required"`
}

type CreatedBy struct {
	Id       string `json:"id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Username string `json:"username" binding:"required" example:"admin.smart"`
}

type GetMaterialsIssuedParams struct {
	paramsDomain.Params
	StartDate *string `json:"start_date" example:"2026-01-01"`
	EndDate   *string `json:"end_date" example:"2026-12-31"`
}

type CreateMaterialIssuedBody struct {
	Description *string `json:"description" example:"Kit de materiales"`
}

type UpdateMaterialIssuedBody struct {
	Description *string `json:"description" example:"Kit de materiales actualizado"`
}

type CreateMaterialIssued struct {
	Id          string  `json:"id"`
	Description *string `json:"description"`
	CreatedBy   string  `json:"created_by"`
}

type UpdateMaterialIssued struct {
	Id          string  `json:"id"`
	Description *string `json:"description"`
}

type DeleteMaterialIssued struct {
	Id        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
}
