package rest

type createWorkshopTypeValidated struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE TALLER"`
	Enable      bool   `json:"enable" example:"true"`
}

type updateWorkshopTypeValidated struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE TALLER"`
	Enable      bool   `json:"enable" example:"true"`
}
