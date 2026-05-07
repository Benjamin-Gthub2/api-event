package rest

type createWorkshopValidated struct {
	TypeId    string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
	Code      *string `json:"code" example:"0001"`
	Capacity  int     `json:"capacity" example:"1"`
	EventId   string  `json:"event_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}

type updateWorkshopValidated struct {
	TypeId    string  `json:"type_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	Name      string  `json:"name" binding:"required" example:"TALLERES - 1ER BLOQUE"`
	Shortname *string `json:"shortname" example:"1ER BLOQUE"`
	Code      *string `json:"code" example:"0001"`
	Capacity  int     `json:"capacity" example:"1"`
}
