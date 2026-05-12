/*
 * File: registrations_certificate_func_pdf_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file content the repository for generate the report.
 *
 * Last Modified: 2026-05-12
 */

package registrations

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	"github.com/Benjamin-Gthub2/api-shared/utils-report"

	registrationsDomain "github.com/Benjamin-Gthub2/api-event/registrations/domain"
)

//go:embed html/template_registration_certificate.html
var TemplateRegistrationCertificate string

//go:embed html/styles.css
var TemplateStyles string

//go:embed html/tejiendocorazones.png
var TemplateImg string

func (c registrationCertificatesReportPdfRepo) GenerateRegistrationReportPdf(
	ctx context.Context,
	registration *registrationsDomain.Registration,
) (fileBin []byte, err error) {
	var errPdf error

	defer logErrorCoreDomain.PanicRecovery(&ctx, &errPdf)
	report, err := utils_report.CreateUtilReport(nil)
	if err != nil {
		return nil, err
	}
	defer report.Clear()
	namePerson := registration.CreatedBy.Person.Names
	lastName := registration.CreatedBy.Person.LastName
	surnameName := registration.CreatedBy.Person.Surname
	arrayName := []*string{namePerson, lastName, surnameName}
	var nameCreatedBy string
	nameCreatedBy, err = report.ConcatStr(arrayName)
	if err != nil {
		return nil, err
	}

	var pathTmpStyle *string
	var pathTmpImg string
	var pathTmpImgAux *string

	fmt.Println(report)
	if LogoMerchantFile == nil {
		pathTmpImgAux, err = report.LoadTextFile(TemplateImg, ".png")
		if err != nil {
			return nil, err
		}
		pathTmpImg = *pathTmpImgAux
	} else {
		storagePath := os.Getenv("STORAGE_PATH")
		url := *LogoMerchantFile.Url
		pathTmpImg = filepath.Join(storagePath, url)
		if _, errFindFile := os.Stat(pathTmpImg); os.IsNotExist(errFindFile) {
			pathTmpImgAux, err = report.LoadTextFile(TemplateImg, ".png")
			if err != nil {
				return nil, err
			}
			pathTmpImg = *pathTmpImgAux
		}
	}

	pathTmpStyle, err = report.LoadTextFile(TemplateStyles, ".css")

	dataReport := dataRegistrationReport{
		PathStyle:                 *pathTmpStyle,
		PathImg:                   pathTmpImg,
		RegistrationConfiguration: *registration,
		NamePerson:                nameCreatedBy,
	}

	var srcTemplateAux *string
	srcTemplateAux, errPdf = report.RenderFileToReport(TemplateRegistration, dataReport)
	if errPdf != nil {
		return nil, errPdf
	}
	processingTemplates := utils_report.ProcessingPdfEntity{
		PathTmpTemplate: srcTemplateAux,
	}

	var pathTmpTemplatePDF *string
	pathTmpTemplatePDF, errPdf = report.GenerateReportPdf(processingTemplates)
	if errPdf != nil {
		return nil, errPdf
	}

	fileBin, err = os.ReadFile(*pathTmpTemplatePDF)
	if err != nil {
		return nil, err
	}

	return fileBin, nil
}
