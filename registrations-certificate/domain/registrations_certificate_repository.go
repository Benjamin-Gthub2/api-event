/*
 * File: registrations_certificate_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2026-05-12
 */

package domain

import (
	"context"

	merchantSharedDomain "github.com/smart0n3/api-shared/merchant-shared/domain"

	requirementsDomain "github.com/smart0n3/api-logistics/requirements/domain"
)

type RegistrationCertificateRepository interface {
	GenerateRegistrationCertificatePdf(ctx context.Context, requirement *requirementsDomain.RequirementById,
		requirementDetails []requirementsDomain.ApprovalPerson, onlyApprovers []requirementsDomain.ApprovalPerson,
		LogoMerchantFile *merchantSharedDomain.MerchantLogo) ([]byte, error)
}
