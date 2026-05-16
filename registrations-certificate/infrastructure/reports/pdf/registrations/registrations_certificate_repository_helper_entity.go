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
	PathStyle                 string                           `json:"path_style"`
	PathImgMain               string                           `json:"path_img_main"`
	PathImgLogo               string                           `json:"path_img_logo"`
	PathImgFirma              string                           `json:"path_img_firma"`
	PathImgAdorno             string                           `json:"path_img_adorno"`
	RegistrationConfiguration registrationsDomain.Registration `json:"registration_configuration"`
	NamePerson                string                           `json:"name_person"`
}
