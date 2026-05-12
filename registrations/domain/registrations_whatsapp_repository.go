/*
 * File: registrations_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository interface for registrations.
 *
 * Last Modified: 2026-05-11
 */

package domain

import "context"

type RegistrationsWhatsAppRepository interface {
	UploadMedia(ctx context.Context, imageBytes []byte) (string, error)
	SendTemplateMessage(ctx context.Context, params SendWhatsAppTemplateParams) error
}

type SendWhatsAppTemplateParams struct {
	To           string
	TemplateName string
	Language     string
	MediaId      string
	Names        string
	EventName    string
}
