/*
 * File: registrations_func_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository functions for registrations (ApisPeru).
 *
 * Last Modified: 2026-05-12
 */

package whatsapp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

// SendImageMessage sends a WhatsApp image message via ApisPeru's /send/image endpoint.
// The QR image is sent as multipart/form-data with a text caption.
func (r registrationsWhatsAppRepository) SendImageMessage(
	ctx context.Context,
	params registrationsDomain.SendWhatsAppImageParams,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	url := fmt.Sprintf("%s/send/image", r.baseURL)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("phone", params.To+"@s.whatsapp.net")
	_ = writer.WriteField("caption", params.Caption)

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", `form-data; name="image"; filename="qr.png"`)
	partHeader.Set("Content-Type", "image/png")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return r.err.Clone().SetFunction("SendImageMessage").SetRaw(err)
	}
	if _, err = io.Copy(part, bytes.NewReader(params.Image)); err != nil {
		return r.err.Clone().SetFunction("SendImageMessage").SetRaw(err)
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &body)
	if err != nil {
		return r.err.Clone().SetFunction("SendImageMessage").SetRaw(err)
	}
	req.Header.Set("Authorization", "Bearer "+r.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if r.deviceId != "" {
		req.Header.Set("device-id", r.deviceId)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return r.err.Clone().SetFunction("SendImageMessage").SetRaw(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		raw, _ := io.ReadAll(resp.Body)
		return r.err.Clone().SetFunction("SendImageMessage").SetRaw(
			fmt.Errorf("apisperu whatsapp send failed (%d): %s", resp.StatusCode, string(raw)),
		)
	}
	return nil
}
