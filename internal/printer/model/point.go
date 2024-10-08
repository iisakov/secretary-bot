package model

type Point struct {
	x Coordinate
	y Coordinate
}

func NewPoint(x, y Coordinate) *Point {
	return &Point{x, y}
}

func (p Point) X() Coordinate {
	return p.x
}

func (p Point) FloatX() float64 {
	return float64(p.x)
}

func (p Point) Y() Coordinate {
	return p.y
}

func (p Point) FloatY() float64 {
	return float64(p.y)
}

func (p *Point) SetX(x Coordinate) *Point {
	p.x = x
	return p
}

func (p *Point) SetY(y Coordinate) *Point {
	p.y = y
	return p
}

func (p *Point) ShiftX(s float64) *Point {
	p.x += Coordinate(s)
	return p
}

func (p *Point) ShiftY(s float64) *Point {
	p.y += Coordinate(s)
	return p
}

func (p Point) Division(d float64) *Point {
	return &Point{p.x / Coordinate(d), p.y / Coordinate(d)}
}
