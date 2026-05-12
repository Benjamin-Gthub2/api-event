/*
 * File: registrations_certificate_repository_helper_entity.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the helper entity for the registrations report.
 *
 * Last Modified: 2026-05-12
 */

package registrations

import (
	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

type dataRegistrationReport struct {
	PathStyle                 string                           `json:"path_style"`
	PathImg                   string                           `json:"path_img"`
	RegistrationConfiguration registrationsDomain.Registration `json:"registration_configuration"`
	NamePerson                string                           `json:"name_person"`
}
