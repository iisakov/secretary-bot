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

func (p *Point) SetX(x Coordinate) {
	p.x = x
}

func (p Point) Y() Coordinate {
	return p.y
}

func (p Point) FloatY() float64 {
	return float64(p.y)
}

func (p *Point) SetY(y Coordinate) {
	p.y = y
}

func (p Point) Division(d float64) *Point {
	return &Point{p.x / Coordinate(d), p.y / Coordinate(d)}
}
