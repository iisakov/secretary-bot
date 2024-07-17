package model

type Printer interface {
	Configurate(nd string, o Options) Printer
	SetColor(c string)
	SetFont(n string)
	AddPage()

	GetFontVariants() *FontVariants
	GetPageSize() *Point

	PrintGrid(x, y Coordinate, opt string)
	PrintLine(l Line) Point
	PrintText(t Text) Point
	PrintCell(c Cell) Point
	PrintTextBR(t Text, maxlen Coordinate) Point

	OutputDoc(domen string) error
}
