package rest

type createRegistrationStatusValidated struct {
	Code        string `json:"code" binding:"required" example:"REGISTERED"`
	Description string `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int    `json:"position" binding:"required" example:"1"`
	Enable      bool   `json:"enable" example:"true"`
}

type updateRegistrationStatusValidated struct {
	Code        string `json:"code" binding:"required" example:"REGISTERED"`
	Description string `json:"description" binding:"required" example:"REGISTRADO"`
	Position    int    `json:"position" binding:"required" example:"1"`
	Enable      bool   `json:"enable" example:"true"`
}
