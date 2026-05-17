/*
 * File: registrations_certificate_repository_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity for the registrations certificate report.
 *
 * Last Modified: 2026-05-16
 */

package registrations

import (
	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type dataRegistrationReport struct {
	StyleCSS                  string                           `json:"style_css"`
	ImgMainB64                string                           `json:"img_main_b64"`
	ImgLogoB64                string                           `json:"img_logo_b64"`
	ImgFirmaB64               string                           `json:"img_firma_b64"`
	ImgAdornoB64              string                           `json:"img_adorno_b64"`
	ImgFirmaAdornoLogoB64     string                           `json:"img_firma_adorno_logo_b64"`
	FontPoppinsRegularB64     string                           `json:"font_poppins_regular_b64"`
	FontPoppinsBoldB64        string                           `json:"font_poppins_bold_b64"`
	RegistrationConfiguration registrationsDomain.Registration `json:"registration_configuration"`
	NamePerson                string                           `json:"name_person"`
}
