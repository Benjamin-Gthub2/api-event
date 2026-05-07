package rest

type createMaterialIssuedValidated struct {
	Description *string `json:"description" example:"Kit de materiales"`
}

type updateMaterialIssuedValidated struct {
	Description *string `json:"description" example:"Kit de materiales actualizado"`
}
