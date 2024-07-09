package model

type Printer interface {
	Configurate(nd string, o Options) Printer
	SetColor(c string)
	SetFont(n string)
	AddPage()

	GetFontVariants() *FontVariants
	GetPageSize() *Point

	PrintGrid(x, y Coordinate, opt string)
	PrintLine(l Line)
	PrintText(t Text)
	PrintTextInCell(t Text, c Cell)

	OutputDoc(domen string) error
}
