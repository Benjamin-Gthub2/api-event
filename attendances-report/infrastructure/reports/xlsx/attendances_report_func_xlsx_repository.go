package xlsx

import (
	"bytes"
	"context"

	"github.com/xuri/excelize/v2"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
	xlsxStyles "github.com/Benjamin-Gthub2/api-event/attendances-report/infrastructure/reports/xlsx/attendances_styles"
)

type attendanceRow struct {
	Number              int
	Event               string
	Workshop            string
	WorkshopType        string
	BeneficiaryFullName string
	TypeDocument        string
	Document            string
	RegisteredBy        string
	CreatedAt           string
}

func (r attendancesReportXlsxRepo) GenerateAttendancesReportXlsx(
	ctx context.Context,
	attendances []attendancesDomain.Attendance,
) (file *bytes.Buffer, err error) {
	newFile := excelize.NewFile()
	defer func() {
		if errClose := newFile.Close(); errClose != nil {
			logErrorCoreDomain.PanicRecovery(&ctx, &errClose)
		}
	}()

	sheetName := "Asistencias"
	index, err := newFile.NewSheet(sheetName)
	if err != nil {
		logErrorCoreDomain.PanicRecovery(&ctx, &err)
		return nil, err
	}

	err = newFile.MergeCell(sheetName, "A1", "I1")
	if err != nil {
		return nil, err
	}
	err = newFile.SetCellValue(sheetName, "A1", "REPORTE DE ASISTENCIAS")
	if err != nil {
		return nil, err
	}
	newFile.SetActiveSheet(index)

	styles, err := xlsxStyles.CreateStyles(newFile)
	if err != nil {
		logErrorCoreDomain.PanicRecovery(&ctx, &err)
		return nil, err
	}

	err = newFile.SetCellStyle(sheetName, "A1", "I1", styles.Colours.GreenBold)
	if err != nil {
		return nil, err
	}
	err = newFile.SetCellStyle(sheetName, "A2", "I2", styles.Colours.BlueBold)
	if err != nil {
		return nil, err
	}

	rowHeaders := []xlsxStyles.RowHeaders{
		{Name: "N°", Width: 6},
		{Name: "EVENTO", Width: 40},
		{Name: "TALLER", Width: 40},
		{Name: "TIPO TALLER", Width: 25},
		{Name: "BENEFICIARIO", Width: 40},
		{Name: "TIPO DOCUMENTO", Width: 20},
		{Name: "DOCUMENTO", Width: 20},
		{Name: "REGISTRADO POR", Width: 25},
		{Name: "FECHA REGISTRO", Width: 22},
	}
	xlsxStyles.SetRowHeaders(rowHeaders, newFile, 2, sheetName)

	valuesFields := []string{
		"Number",
		"Event",
		"Workshop",
		"WorkshopType",
		"BeneficiaryFullName",
		"TypeDocument",
		"Document",
		"RegisteredBy",
		"CreatedAt",
	}
	rows := buildRows(attendances)
	xlsxStyles.SetCellValueRows(newFile, rows, valuesFields, 3, sheetName)

	err = newFile.DeleteSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	file, err = newFile.WriteToBuffer()
	if err != nil {
		logErrorCoreDomain.PanicRecovery(&ctx, &err)
		return nil, err
	}
	return file, nil
}

func buildRows(attendances []attendancesDomain.Attendance) []attendanceRow {
	rows := make([]attendanceRow, 0, len(attendances))
	for i, a := range attendances {
		fullName := a.Beneficiary.Names + " " + a.Beneficiary.Surname
		if a.Beneficiary.LastName != nil && *a.Beneficiary.LastName != "" {
			fullName += " " + *a.Beneficiary.LastName
		}

		createdAt := ""
		if a.CreatedAt != nil {
			createdAt = a.CreatedAt.Format("02/01/2006 15:04")
		}

		rows = append(rows, attendanceRow{
			Number:              i + 1,
			Event:               a.Workshop.Event.Name,
			Workshop:            a.Workshop.Name,
			WorkshopType:        a.Workshop.WorkshopType.Description,
			BeneficiaryFullName: fullName,
			TypeDocument:        a.Beneficiary.TypeDocument.AbbreviatedDescription,
			Document:            a.Beneficiary.Document,
			RegisteredBy:        a.CreatedBy.Username,
			CreatedAt:           createdAt,
		})
	}
	return rows
}
