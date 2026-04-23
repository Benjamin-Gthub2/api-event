package rest

type createEventTypeValidated struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE EVENTO"`
	Enable      bool   `json:"enable" example:"true"`
}

type updateEventTypeValidated struct {
	Code        string `json:"code" binding:"required" example:"0001"`
	Description string `json:"description" binding:"required" example:"TIPO DE EVENTO"`
	Enable      bool   `json:"enable" example:"true"`
}
