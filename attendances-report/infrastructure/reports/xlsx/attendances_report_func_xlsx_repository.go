package xlsx

import (
	"bytes"
	"context"
	"strings"

	"github.com/xuri/excelize/v2"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	"github.com/Benjamin-Gthub2/api-event/attendances-report/domain"
	xlsxStyles "github.com/Benjamin-Gthub2/api-event/attendances-report/infrastructure/reports/xlsx/attendances_styles"
	attendancesDomain "github.com/Benjamin-Gthub2/api-event/attendances/domain"
)

type attendanceRow struct {
	Number              int
	Event               string
	Workshop            string
	BeneficiaryFullName string
	TypeDocument        string
	Document            string
	RegisteredBy        string
	CreatedAt           string
}

func (r attendancesReportXlsxRepo) GenerateAttendancesReportXlsx(
	ctx context.Context,
	attendances []attendancesDomain.Attendance,
	displayFilters domain.AttendancesReportDisplayFilters,
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
	newFile.SetActiveSheet(index)

	styles, err := xlsxStyles.CreateStyles(newFile)
	if err != nil {
		logErrorCoreDomain.PanicRecovery(&ctx, &err)
		return nil, err
	}

	// Row 1: titulo principal (A1:H1)
	if err = newFile.MergeCell(sheetName, "A1", "H1"); err != nil {
		return nil, err
	}
	if err = newFile.SetCellValue(sheetName, "A1", "REPORTE DE ASISTENCIAS"); err != nil {
		return nil, err
	}
	if err = newFile.SetCellStyle(sheetName, "A1", "H1", styles.Colours.GreenBold); err != nil {
		return nil, err
	}

	// Row 2: filtros aplicados (A2:H2)
	if err = newFile.MergeCell(sheetName, "A2", "H2"); err != nil {
		return nil, err
	}
	if err = newFile.SetCellValue(sheetName, "A2", buildSubtitle(displayFilters)); err != nil {
		return nil, err
	}
	if err = newFile.SetCellStyle(sheetName, "A2", "H2", styles.Colours.BlueBold); err != nil {
		return nil, err
	}

	// Row 3: cabeceras de columnas
	rowHeaders := []xlsxStyles.RowHeaders{
		{Name: "N°", Width: 6},
		{Name: "EVENTO", Width: 40},
		{Name: "TALLER", Width: 40},
		{Name: "BENEFICIARIO", Width: 40},
		{Name: "TIPO DOCUMENTO", Width: 20},
		{Name: "DOCUMENTO", Width: 20},
		{Name: "FECHA ASISTENCIA", Width: 22},
		{Name: "REGISTRADO POR", Width: 25},
	}
	xlsxStyles.SetRowHeaders(rowHeaders, newFile, 3, sheetName)
	if err = newFile.SetCellStyle(sheetName, "A3", "H3", styles.Colours.BlueBold); err != nil {
		return nil, err
	}

	// Rows 4+: datos
	valuesFields := []string{
		"Number",
		"Event",
		"Workshop",
		"BeneficiaryFullName",
		"TypeDocument",
		"Document",
		"CreatedAt",
		"RegisteredBy",
	}
	rows := buildRows(attendances)
	xlsxStyles.SetCellValueRows(newFile, rows, valuesFields, 4, sheetName)

	if err = newFile.DeleteSheet("Sheet1"); err != nil {
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
			BeneficiaryFullName: fullName,
			TypeDocument:        a.Beneficiary.TypeDocument.AbbreviatedDescription,
			Document:            a.Beneficiary.Document,
			RegisteredBy:        a.CreatedBy.Username,
			CreatedAt:           createdAt,
		})
	}
	return rows
}

func buildSubtitle(filters domain.AttendancesReportDisplayFilters) string {
	parts := []string{}
	if filters.EventName != "" {
		parts = append(parts, "Evento: "+filters.EventName)
	}
	if filters.WorkshopName != "" {
		parts = append(parts, "Taller: "+filters.WorkshopName)
	}
	if filters.BeneficiaryName != "" {
		parts = append(parts, "Beneficiario: "+filters.BeneficiaryName)
	}
	if len(parts) == 0 {
		return "Todos los registros"
	}
	return strings.Join(parts, " | ")
}
