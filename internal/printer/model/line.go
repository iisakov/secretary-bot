package model

import "math"

type Line struct {
	Start  Point
	Finish Point
	Width  Coordinate
}

func NewHorisontLine(hight, startX, finishX Coordinate, width Coordinate) *Line {
	return &Line{Point{startX, hight}, Point{finishX, hight}, width}
}

func NewVerticalLine(shift, startY, finishY Coordinate, width Coordinate) *Line {
	return &Line{Point{shift, startY}, Point{shift, finishY}, width}
}

func (l Line) Len() Coordinate {
	result := math.Sqrt(math.Pow(l.Finish.FloatX()-l.Start.FloatX(), 2) + math.Pow(l.Finish.FloatY()-l.Start.FloatY(), 2))
	return Coordinate(result)
}
