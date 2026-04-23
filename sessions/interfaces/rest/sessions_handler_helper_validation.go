package rest

type createSessionValidated struct {
	WorkshopId string `json:"workshop_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	StartDate  string `json:"start_date" binding:"required" example:"2026-04-21 09:00:00"`
	EndDate    string `json:"end_date" binding:"required" example:"2026-04-21 11:00:00"`
}

type updateSessionValidated struct {
	StartDate string `json:"start_date" binding:"required" example:"2026-04-21 09:00:00"`
	EndDate   string `json:"end_date" binding:"required" example:"2026-04-21 11:00:00"`
}
