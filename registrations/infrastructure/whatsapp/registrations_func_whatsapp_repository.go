/*
 * File: registrations_func_whatsapp_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the WhatsApp repository functions for registrations.
 *
 * Last Modified: 2026-05-11
 */

package whatsapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

// UploadMedia uploads a PNG image to WhatsApp Media API and returns the media ID.
func (r registrationsWhatsAppRepository) UploadMedia(
	ctx context.Context,
	imageBytes []byte,
) (
	mediaId string,
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	url := fmt.Sprintf("%s/%s/%s/media", r.baseURL, r.apiVersion, r.phoneNumberId)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("messaging_product", "whatsapp")

	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", `form-data; name="file"; filename="qr.png"`)
	partHeader.Set("Content-Type", "image/png")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(err)
	}
	if _, err = io.Copy(part, bytes.NewReader(imageBytes)); err != nil {
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(err)
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &body)
	if err != nil {
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(err)
	}
	req.Header.Set("Authorization", "Bearer "+r.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body)
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(
			fmt.Errorf("whatsapp media upload failed (%d): %s", resp.StatusCode, string(raw)),
		)
	}

	var result struct {
		Id string `json:"id"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", r.err.Clone().SetFunction("UploadMedia").SetRaw(err)
	}
	return result.Id, nil
}

// SendTemplateMessage sends a WhatsApp template message with an image header.
//
// The template must be pre-approved in Meta Business Manager.
// Expected template structure:
//
//	Header  : IMAGE  ({{1}} = QR image)
//	Body    : "Hola {{1}}, tu código QR para el evento *{{2}}* está listo.
//	           Preséntalo el día del evento para registrar tu asistencia."
//	Footer  : "Smart Cities Peru"
func (r registrationsWhatsAppRepository) SendTemplateMessage(
	ctx context.Context,
	params registrationsDomain.SendWhatsAppTemplateParams,
) (
	err error,
) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	url := fmt.Sprintf("%s/%s/%s/messages", r.baseURL, r.apiVersion, r.phoneNumberId)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                params.To,
		"type":              "template",
		"template": map[string]interface{}{
			"name":     params.TemplateName,
			"language": map[string]string{"code": params.Language},
			"components": []map[string]interface{}{
				{
					"type": "header",
					"parameters": []map[string]interface{}{
						{
							"type":  "image",
							"image": map[string]string{"id": params.MediaId},
						},
					},
				},
				{
					"type": "body",
					"parameters": []map[string]interface{}{
						{"type": "text", "text": params.Names},
						{"type": "text", "text": params.EventName},
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return r.err.Clone().SetFunction("SendTemplateMessage").SetRaw(err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return r.err.Clone().SetFunction("SendTemplateMessage").SetRaw(err)
	}
	req.Header.Set("Authorization", "Bearer "+r.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return r.err.Clone().SetFunction("SendTemplateMessage").SetRaw(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		raw, _ := io.ReadAll(resp.Body)
		return r.err.Clone().SetFunction("SendTemplateMessage").SetRaw(
			fmt.Errorf("whatsapp send failed (%d): %s", resp.StatusCode, string(raw)),
		)
	}
	return nil
}
