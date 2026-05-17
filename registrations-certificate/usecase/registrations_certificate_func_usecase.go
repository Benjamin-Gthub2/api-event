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
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func sanitizeFileName(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

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
	fullName := strings.Join(nameParts, " ")
	fileName = strings.ReplaceAll(sanitizeFileName(strings.ToUpper(fullName)), " ", "_") + "_certificado.pdf"

	// Verificar caché R2 primero para evitar regenerar un PDF ya generado.
	cached, err := u.registrationCertificateStorageRepository.GetCertificate(ctx, registrationId)
	if err != nil {
		return nil, "", u.err.Clone().SetFunction("GenerateRegistrationsCertificatePdf").SetRaw(err)
	}
	if cached != nil {
		return cached, fileName, nil
	}

	pdfBytes, err = u.registrationCertificateRepository.GenerateRegistrationCertificatePdf(ctx, registration)
	if err != nil {
		return nil, "", u.err.Clone().SetFunction("GenerateRegistrationsCertificatePdf").SetRaw(err)
	}

	// Upload asynchronously so the caller is not blocked by R2 latency.
	go func() {
		_ = u.registrationCertificateStorageRepository.UploadCertificate(context.Background(), registrationId, pdfBytes)
	}()

	return pdfBytes, fileName, nil
}
