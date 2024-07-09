package model

type Text struct {
	Name        string
	FVName      string
	Orientation Orientation
	Text        string
	Line        string
}

type Orientation struct {
	Start  Point
	Align  string
	Space  Space
	Border string
}

type Space [4]Coordinate

func (s Space) Left() Coordinate {
	return s[0]
}

func (s Space) Top() Coordinate {
	return s[1]
}

func (s Space) Right() Coordinate {
	return s[2]
}

func (s Space) Bottom() Coordinate {
	return s[3]
}
