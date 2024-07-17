package model

type Text struct {
	Name        string
	FVName      string
	Orientation Orientation
	Text        string
	Line        string
}

type Orientation struct {
	Start   Point
	Align   string
	Space   Space
	Border  string
	Padding Coordinate
	Indent  Indent
}

type Indent struct {
	Indent   Coordinate
	NumLines uint
}

type TextBuilderable interface {
	Name(val string) TextBuilderable
	FVName(val string) TextBuilderable
	Orientation(val Orientation) TextBuilderable
	Text(val string) TextBuilderable
	Line(val string) TextBuilderable

	Build() Text
}

type TextBuilder struct {
	name        string
	fvName      string
	orientation Orientation
	text        string
	line        string
}

func NewTextBuilder() TextBuilderable {
	return TextBuilder{}
}

func (tb TextBuilder) Name(val string) TextBuilderable {
	tb.name = val
	return tb
}
func (tb TextBuilder) FVName(val string) TextBuilderable {
	tb.fvName = val
	return tb
}
func (tb TextBuilder) Orientation(val Orientation) TextBuilderable {
	tb.orientation = val
	return tb
}
func (tb TextBuilder) Text(val string) TextBuilderable {
	tb.text = val
	return tb
}
func (tb TextBuilder) Line(val string) TextBuilderable {
	tb.line = val
	return tb
}

func (tb TextBuilder) Build() Text {
	return Text{
		Name:        tb.name,
		FVName:      tb.fvName,
		Orientation: tb.orientation,
		Text:        tb.text,
		Line:        tb.line,
	}
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
