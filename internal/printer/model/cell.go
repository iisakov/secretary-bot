package model

type Cell struct {
	Border    string
	Color     string
	Lines     [4]Line
	Text      Text
	Indent    Indent
	TextAlign string

	tlPoint Point
	trPoint Point
	brPoint Point
	blPoint Point
}

type Indent struct {
	Indent   Coordinate
	NumLines uint
}

func NewCell(tlPoint, brPoint Point, width Coordinate, border, color string) *Cell {
	c := &Cell{Lines: [4]Line{
		*NewVerticalLine(tlPoint.X()-width/2+1, tlPoint.Y()-width+1, brPoint.Y()+width-1, width),
		*NewHorisontLine(tlPoint.Y()-width/2+1, tlPoint.X()-width+1, brPoint.X()+width-1, width),
		*NewVerticalLine(brPoint.X()+width/2-1, tlPoint.Y()-width+1, brPoint.Y()+width-1, width),
		*NewHorisontLine(brPoint.Y()+width/2-1, brPoint.X()+width-1, tlPoint.X()-width+1, width),
	}}

	c.Border = border
	c.Color = color
	c.tlPoint = Point{tlPoint.X(), tlPoint.Y()}
	c.trPoint = Point{brPoint.X(), tlPoint.Y()}
	c.brPoint = Point{brPoint.X(), brPoint.Y()}
	c.blPoint = Point{tlPoint.X(), brPoint.Y()}

	return c
}

func (c *Cell) AddText(t Text, i Indent, ta string) *Cell {
	c.Text = t
	c.Indent = i
	c.TextAlign = ta
	return c
}

func (c Cell) BAll() *[4]Line {
	return &c.Lines
}

func (c Cell) BLeft() *Line {
	return &c.Lines[0]
}

func (c Cell) BTop() *Line {
	return &c.Lines[1]
}

func (c Cell) BRight() *Line {
	return &c.Lines[2]
}

func (c Cell) BBottom() *Line {
	return &c.Lines[3]
}

func (c Cell) TL() *Point {
	return &c.tlPoint
}

func (c Cell) TR() *Point {
	return &c.trPoint
}

func (c Cell) BR() *Point {
	return &c.brPoint
}

func (c Cell) BL() *Point {
	return &c.blPoint
}
