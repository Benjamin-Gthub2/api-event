/*
 * File: registrations_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository for registrations (ApisPeru).
 *
 * Last Modified: 2026-05-12
 */

package whatsapp

import (
	"net/http"
	"os"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsWhatsAppRepository struct {
	token      string
	deviceId   string
	baseURL    string
	httpClient *http.Client
	err        *errDomain.SmartError
}

func NewRegistrationsWhatsAppRepository() registrationsDomain.RegistrationsWhatsAppRepository {
	return &registrationsWhatsAppRepository{
		token:      os.Getenv("APISPERU_WHATSAPP_TOKEN"),
		deviceId:   os.Getenv("APISPERU_WHATSAPP_DEVICE_ID"),
		baseURL:    "https://whatsapp.apisperu.com/api/v1",
		httpClient: &http.Client{},
		err:        errDomain.NewErr().SetLayer(errDomain.Interface),
	}
}
