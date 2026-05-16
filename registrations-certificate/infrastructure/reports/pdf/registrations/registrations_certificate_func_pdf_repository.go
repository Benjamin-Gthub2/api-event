/*
 * File: registrations_certificate_func_pdf_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file generates the certificate PDF using wkhtmltopdf.
 * Uses html/template to render the HTML and calls wkhtmltopdf via os/exec.
 * All temp files are cleaned up in a deferred call to avoid disk leaks on Railway.
 *
 * Last Modified: 2026-05-16
 */

package registrations

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

//go:embed html/template_registration_certificate.html
var TemplateRegistrationCertificate string

//go:embed html/styles.css
var TemplateStyles string

//go:embed html/tejiendocorazones.png
var TemplateImgMain []byte

//go:embed html/logo.png
var TemplateImgLogo []byte

//go:embed html/firma.png
var TemplateImgFirma []byte

//go:embed html/adorno.png
var TemplateImgAdorno []byte

//go:embed html/Poppins-Regular.ttf
var FontPoppinsRegular []byte

//go:embed html/Poppins-Bold.ttf
var FontPoppinsBold []byte

//go:embed html/frma-adorno-y-logo.png
var TemplateImgFirmaAdornoLogo []byte

func (c registrationCertificatesReportPdfRepo) GenerateRegistrationCertificatePdf(
	ctx context.Context,
	registration *registrationsDomain.Registration,
) (fileBin []byte, err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	names := registration.Beneficiary.Names
	surname := registration.Beneficiary.Surname
	nameParts := []string{names, surname}
	if registration.Beneficiary.LastName != nil {
		nameParts = append(nameParts, *registration.Beneficiary.LastName)
	}
	fullName := strings.Join(nameParts, " ")

	tmpDir, err := os.MkdirTemp("", "cert-*")
	if err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}
	defer os.RemoveAll(tmpDir)

	htmlPath := filepath.Join(tmpDir, "certificate.html")
	pdfPath := filepath.Join(tmpDir, "certificate.pdf")

	tmpl, err := template.New("certificate").Parse(TemplateRegistrationCertificate)
	if err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}

	data := dataRegistrationReport{
		StyleCSS:                  TemplateStyles,
		ImgMainB64:                base64.StdEncoding.EncodeToString(TemplateImgMain),
		ImgLogoB64:                base64.StdEncoding.EncodeToString(TemplateImgLogo),
		ImgFirmaB64:               base64.StdEncoding.EncodeToString(TemplateImgFirma),
		ImgAdornoB64:              base64.StdEncoding.EncodeToString(TemplateImgAdorno),
		ImgFirmaAdornoLogoB64:     base64.StdEncoding.EncodeToString(TemplateImgFirmaAdornoLogo),
		FontPoppinsRegularB64:     base64.StdEncoding.EncodeToString(FontPoppinsRegular),
		FontPoppinsBoldB64:        base64.StdEncoding.EncodeToString(FontPoppinsBold),
		RegistrationConfiguration: *registration,
		NamePerson:                fullName,
	}

	var htmlBuf bytes.Buffer
	if err = tmpl.Execute(&htmlBuf, data); err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}
	if err = os.WriteFile(htmlPath, htmlBuf.Bytes(), 0644); err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}

	cmdCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	wkhtmltopdfBin := "wkhtmltopdf"
	if customPath := os.Getenv("WKHTMLTOPDF_PATH"); customPath != "" {
		wkhtmltopdfBin = filepath.Join(customPath, "wkhtmltopdf")
	}

	cmd := exec.CommandContext(cmdCtx, wkhtmltopdfBin,
		"--orientation", "Landscape",
		"--page-size", "A4",
		"--margin-top", "0",
		"--margin-right", "0",
		"--margin-bottom", "0",
		"--margin-left", "0",
		"--disable-smart-shrinking",
		"--print-media-type",
		"--quiet",
		"--no-stop-slow-scripts",
		htmlPath,
		pdfPath,
	)
	if output, errExec := cmd.CombinedOutput(); errExec != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").
			SetRaw(fmt.Errorf("wkhtmltopdf: %w, output: %s", errExec, output))
	}

	fileBin, err = os.ReadFile(pdfPath)
	if err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}

	return fileBin, nil
}
