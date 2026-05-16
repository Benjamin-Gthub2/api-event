/*
 * File: registrations_certificate_func_usecase.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the use case functions for registrations certificate.
 *
 * Last Modified: 2026-05-16
 */

package usecase

import (
	"context"
	"strings"
)

func (u registrationsCertificateUseCase) GenerateRegistrationsCertificatePdf(
	ctx context.Context,
	registrationId string,
) (pdfBytes []byte, fileName string, err error) {
	registration, err := u.registrationsRepository.GetRegistrationById(ctx, registrationId)
	if err != nil {
		return nil, "", u.err.Clone().SetFunction("GenerateRegistrationsCertificatePdf").SetRaw(err)
	}

	names := registration.Beneficiary.Names
	surname := registration.Beneficiary.Surname
	nameParts := []string{names, surname}
	if registration.Beneficiary.LastName != nil {
		nameParts = append(nameParts, *registration.Beneficiary.LastName)
	}
	fullName := strings.Join(nameParts, "_")
	fileName = strings.ReplaceAll(fullName, " ", "_") + "_certificado.pdf"

	pdfBytes, err = u.registrationCertificateRepository.GenerateRegistrationCertificatePdf(ctx, registration)
	if err != nil {
		return nil, "", u.err.Clone().SetFunction("GenerateRegistrationsCertificatePdf").SetRaw(err)
	}

	return pdfBytes, fileName, nil
}
