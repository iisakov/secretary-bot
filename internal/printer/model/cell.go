package model

type Cell struct {
	Lines [4]Line
}

func NewCell(tlPoint, brPoint Point, width Coordinate) *Cell {

	return &Cell{Lines: [4]Line{
		*NewVerticalLine(tlPoint.X()-width/2, tlPoint.Y()-width, brPoint.Y()+width, width),
		*NewHorisontLine(tlPoint.Y()-width/2, tlPoint.X()-width, brPoint.X()+width, width),
		*NewVerticalLine(brPoint.X()+width/2, tlPoint.Y()-width, brPoint.Y()+width, width),
		*NewHorisontLine(brPoint.Y()+width/2, brPoint.X()+width, tlPoint.X()-width, width),
	}}
}

func (c Cell) Border() *[4]Line {
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
