package domain

import (
	"net/http"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

const (
	ErrWorkshopSpeakerNotFoundCode = "ERR_WORKSHOP_SPEAKER_NOT_FOUND"
)

var (
	ErrWorkshopSpeakerNotFound = errDomain.NewErr().
					SetCode(ErrWorkshopSpeakerNotFoundCode).
					SetDescription("WORKSHOP SPEAKER NOT FOUND").
					SetLevel(errDomain.LevelError).
					SetHttpStatus(http.StatusNotFound).
					SetLayer(errDomain.UseCase).
					SetFunction("GetWorkshopSpeakerById")
)
