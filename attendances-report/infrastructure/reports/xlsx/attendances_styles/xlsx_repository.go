package attendances_styles

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

var (
	colorBlueHex   = "#7CC0F9"
	colorGreenHex  = "AEDFBB"
	colorLilacHex  = "#C0ABF3"
	colorOrangeHex = "#FFC485"
)

func CreateStyles(newFile *excelize.File) (styles StylesGenerals, err error) {
	var colorStyles, colorStylesDistinct []int
	var headerStyle, center, styleNumber, styleNumberSmall, styleDateDMY, styleDate, styleDateSmall, styleTime,
		styleText, rotation90, sizeSmall, styleTextSmall, styleAmount int
	var formatText = "@"
	alignment := &excelize.Alignment{
		Horizontal: "center", Vertical: "center",
	}

	alignmentAmount := &excelize.Alignment{
		Horizontal: "right",
	}

	alignmentMonths := &excelize.Alignment{
		TextRotation: 90, Horizontal: "center", Vertical: "center",
	}

	font := &excelize.Font{
		Bold: true,
	}

	fontSizeSmall := &excelize.Font{
		Size: 9,
	}

	headerStyle, err = newFile.NewStyle(&excelize.Style{
		Alignment:    alignment,
		Font:         font,
		CustomNumFmt: &formatText,
	})
	if err != nil {
		return styles, err
	}

	center, err = newFile.NewStyle(&excelize.Style{
		Alignment:    alignment,
		CustomNumFmt: &formatText,
	})
	if err != nil {
		return styles, err
	}

	styleNumber, err = newFile.NewStyle(&excelize.Style{NumFmt: 0})
	if err != nil {
		return styles, err
	}

	styleNumberSmall, err = newFile.NewStyle(&excelize.Style{
		Font:   fontSizeSmall,
		NumFmt: 0})
	if err != nil {
		return styles, err
	}

	styleDateDMY, err = newFile.NewStyle(&excelize.Style{
		NumFmt: 14})
	if err != nil {
		return styles, err
	}

	styleDate, err = newFile.NewStyle(&excelize.Style{
		NumFmt: 15})
	if err != nil {
		return styles, err
	}

	styleDateSmall, err = newFile.NewStyle(&excelize.Style{
		Font:   fontSizeSmall,
		NumFmt: 15})
	if err != nil {
		return styles, err
	}

	styleTime, err = newFile.NewStyle(&excelize.Style{NumFmt: 20})
	if err != nil {
		return styles, err
	}

	styleText, err = newFile.NewStyle(&excelize.Style{
		CustomNumFmt: &formatText,
	})
	if err != nil {
		return styles, err
	}

	styleTextSmall, err = newFile.NewStyle(&excelize.Style{
		CustomNumFmt: &formatText,
		Font:         fontSizeSmall,
	})
	if err != nil {
		return styles, err
	}

	rotation90, err = newFile.NewStyle(&excelize.Style{
		Alignment: alignmentMonths,
		Font:      font,
	})
	if err != nil {
		return styles, err
	}

	colorStyles, err = ColorStyles(newFile)
	if err != nil {
		return styles, err
	}
	colorStylesDistinct, err = ColorStylesDistinct(newFile)
	if err != nil {
		return styles, err
	}

	styleAmount, err = newFile.NewStyle(&excelize.Style{
		Alignment: alignmentAmount,
	})
	if err != nil {
		return styles, err
	}

	styles = StylesGenerals{
		Rotate90: rotation90, CenterBold: headerStyle, Center: center,
		Number: styleNumber, Date: styleDate,
		DateDMY: styleDateDMY, DateFormatSmall: styleDateSmall,
		NumberFormatSmall: styleNumberSmall, TextSmall: styleTextSmall,
		SizeSmall: sizeSmall, Hour: styleTime, Text: styleText, Amount: styleAmount,
		Colours: Colours{
			Blue:       colorStyles[0],
			GrayBold:   colorStyles[1],
			BlueBold:   colorStyles[2],
			Green:      colorStyles[3],
			Orange:     colorStyles[4],
			Lilac:      colorStyles[5],
			GreenBold:  colorStyles[6],
			OrangeBold: colorStyles[7],
			LilacBold:  colorStyles[8],
		},
		ColoursValues: colorStylesDistinct,
	}
	return styles, nil
}

func ColorStyles(newFile *excelize.File) ([]int, error) {
	var err error
	var colorBlue, colorBlueBold, colorGrayBold, colorGreen, colorLilac, colorOrange, colorGreenBold, colorOrangeBold,
		colorLilacBold int
	alignment := &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		WrapText:   true,
	}

	font := &excelize.Font{
		Bold: true,
	}

	colorBlue, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorBlueHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorBlueBold, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorBlueHex},
			Pattern: 1,
		},
		Font:      font,
		Alignment: alignment,
	})
	if err != nil {
		return []int{}, err
	}

	colorGrayBold, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C2C2C2"},
			Pattern: 1,
		},
		Font:      font,
		Alignment: alignment,
	})
	if err != nil {
		return []int{}, err
	}

	colorGreen, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorGreenHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorOrange, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorOrangeHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorLilac, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorLilacHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorGreenBold, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorGreenHex},
			Pattern: 1,
		},
		Font:      font,
		Alignment: alignment,
	})
	if err != nil {
		return []int{}, err
	}

	colorOrangeBold, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorOrangeHex},
			Pattern: 1,
		},
		Font:      font,
		Alignment: alignment,
	})
	if err != nil {
		return []int{}, err
	}

	colorLilacBold, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorLilacHex},
			Pattern: 1,
		},
		Font:      font,
		Alignment: alignment,
	})
	if err != nil {
		return []int{}, err
	}

	return []int{colorBlue, colorGrayBold, colorBlueBold, colorGreen, colorOrange, colorLilac, colorGreenBold,
		colorOrangeBold, colorLilacBold}, nil
}

func ColorStylesDistinct(newFile *excelize.File) ([]int, error) {
	var err error
	var colorBlue, colorGreen, colorLilac, colorOrange int
	var colorPink, colorGreenOil, colorLightBlue, colorGreenBlack, colorRed int

	colorBlue, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorBlueHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorGreen, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorGreenHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorOrange, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{colorOrangeHex},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorLilac, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#C0ABF3"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorPink, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#F789F2"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorGreenOil, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#ACAF7D"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorLightBlue, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#ACFFE7"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorGreenBlack, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#1DD0B9"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	colorRed, err = newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FF5777"},
			Pattern: 1,
		},
	})
	if err != nil {
		return []int{}, err
	}

	return []int{colorGreen, colorOrange, colorLilac,
		colorPink, colorGreenOil, colorLightBlue, colorGreenBlack, colorRed, colorBlue}, nil
}

func SetRowHeaders(rowHeaders []RowHeaders, file *excelize.File, posInit int, sheetName string) {
	letterColumnsAux := LetterColumns()
	positionAux := strconv.Itoa(posInit)
	var err error
	for index, row := range rowHeaders {
		err = file.SetCellValue(sheetName, letterColumnsAux[index]+positionAux, row.Name)
		if err != nil {
			return
		}
		err = file.SetColWidth(sheetName, letterColumnsAux[index], letterColumnsAux[index], row.Width)
		if err != nil {
			return
		}
	}
}

func LetterColumns() []string {
	length := 50
	columns := make([]string, length)
	for i := 0; i < 26; i++ {
		columns[i] = string(rune('A' + i))
	}
	for i := 26; i < length; i++ {
		div := i / 26
		mod := i % 26
		columns[i] = string(rune('A'+div-1)) + string(rune('A'+mod))
	}
	return columns
}

func SetCellValueRows(file *excelize.File, items interface{}, valuesRows []string, posInit int, sheetName string) {
	letterColumnsAux := LetterColumns()
	itemsSlice := reflect.ValueOf(items)
	for indexSlice := 0; indexSlice < itemsSlice.Len(); indexSlice++ {
		item := itemsSlice.Index(indexSlice).Interface()
		auxItem := reflect.ValueOf(item)
		for index, fieldName := range valuesRows {
			fieldValue := GetDynamicFieldValue(fieldName, auxItem)
			positionAux := strconv.Itoa(posInit)
			err := file.SetCellValue(sheetName, letterColumnsAux[index]+positionAux, fieldValue)
			if err != nil {
				return
			}
		}
		posInit++
	}
}

func GetDynamicFieldValue(fieldName string, auxItem reflect.Value) interface{} {
	names := strings.Split(fieldName, ".")
	fieldValue := auxItem
	var value interface{}
	for _, name := range names {
		if fieldValue.Kind() != reflect.Invalid {
			fieldValue = fieldValue.FieldByName(name)
			value = UnpackPointerOrValue(fieldValue.Interface())
		} else {
			return reflect.Value{}
		}
	}
	return value
}

func UnpackPointerOrValue(data interface{}) interface{} {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		if !val.IsNil() {
			return val.Elem().Interface()
		}
		return ""
	}
	return data
}
