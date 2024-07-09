package modal

type FontsOptions struct {
	FamilyStr string
	StyleStr  string
	FileStr   string
}

type Bluprint struct {
	NameDoc      string
	BaseFontSize float64

	Fonts []FontsOptions

	OrientationStr string
	UnitStr        string
	SizeStr        string
	FontDirStr     string
}
