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
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	imgMainPath := filepath.Join(tmpDir, "main.png")
	imgLogoPath := filepath.Join(tmpDir, "logo.png")
	imgFirmaPath := filepath.Join(tmpDir, "firma.png")
	imgAdornoPath := filepath.Join(tmpDir, "adorno.png")
	stylePath := filepath.Join(tmpDir, "styles.css")
	htmlPath := filepath.Join(tmpDir, "certificate.html")
	pdfPath := filepath.Join(tmpDir, "certificate.pdf")

	for path, data := range map[string][]byte{
		imgMainPath:   TemplateImgMain,
		imgLogoPath:   TemplateImgLogo,
		imgFirmaPath:  TemplateImgFirma,
		imgAdornoPath: TemplateImgAdorno,
		stylePath:     []byte(TemplateStyles),
	} {
		if err = os.WriteFile(path, data, 0644); err != nil {
			return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
		}
	}

	tmpl, err := template.New("certificate").Parse(TemplateRegistrationCertificate)
	if err != nil {
		return nil, c.err.Clone().SetFunction("GenerateRegistrationCertificatePdf").SetRaw(err)
	}

	data := dataRegistrationReport{
		PathStyle:                 stylePath,
		PathImgMain:               imgMainPath,
		PathImgLogo:               imgLogoPath,
		PathImgFirma:              imgFirmaPath,
		PathImgAdorno:             imgAdornoPath,
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
