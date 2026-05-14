package attendances_styles

type StylesGenerals struct {
	CenterBold        int
	GrayBold          int
	Number            int
	NumberFormatSmall int
	Date              int
	DateDMY           int
	DateFormatSmall   int
	Text              int
	TextSmall         int
	Rotate90          int
	SizeSmall         int
	Hour              int
	Amount            int
	Colours           Colours
	ColoursValues     []int
	Center            int
}

type Colours struct {
	Blue       int
	GrayBold   int
	BlueBold   int
	Green      int
	Orange     int
	Lilac      int
	GreenBold  int
	OrangeBold int
	LilacBold  int
}

type RowHeaders struct {
	Name   string
	Hex    string
	Values []string
	Width  float64
}
