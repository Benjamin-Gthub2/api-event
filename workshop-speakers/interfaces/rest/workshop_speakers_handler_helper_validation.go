package rest

type createWorkshopSpeakerValidated struct {
	WorkshopId string `json:"workshop_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
	SpeakerId  string `json:"speaker_id" binding:"required" example:"739bbbc9-7e93-11ee-89fd-0242ac113422"`
}
