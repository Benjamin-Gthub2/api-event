/*
 * File: registrations_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository for registrations.
 *
 * Last Modified: 2026-05-11
 */

package whatsapp

import (
	"net/http"
	"os"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type registrationsWhatsAppRepository struct {
	token         string
	phoneNumberId string
	apiVersion    string
	baseURL       string
	httpClient    *http.Client
	err           *errDomain.SmartError
}

func NewRegistrationsWhatsAppRepository() registrationsDomain.RegistrationsWhatsAppRepository {
	return &registrationsWhatsAppRepository{
		token:         os.Getenv("WHATSAPP_TOKEN"),
		phoneNumberId: os.Getenv("WHATSAPP_PHONE_NUMBER_ID"),
		apiVersion:    "v19.0",
		baseURL:       "https://graph.facebook.com",
		httpClient:    &http.Client{},
		err:           errDomain.NewErr().SetLayer(errDomain.Interface),
	}
}
