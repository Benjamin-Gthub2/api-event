/*
 * File: registrations_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository interface for registrations.
 *
 * Last Modified: 2026-05-12
 */

package domain

import "context"

type RegistrationsWhatsAppRepository interface {
	SendImageMessage(ctx context.Context, params SendWhatsAppImageParams) error
}

type SendWhatsAppImageParams struct {
	// To is the destination phone number with country code (e.g. 51987654321).
	// The @s.whatsapp.net suffix is added automatically by the repository.
	To      string
	Caption string
	Image   []byte
}
