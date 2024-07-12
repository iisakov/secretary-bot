package pdfer

import (
	"secretary/internal/printer/model"
	"strings"
)

func (p PDFer) PrintText(t model.Text) model.Point {
	p.SetFont(t.FVName)

	ancor := model.NewPoint(t.Orientation.Start.X(), t.Orientation.Start.Y())
	fontDescriptor := p.GetFontDesc(t.FVName)
	fontSize, _ := p.Pdf.GetFontSize()
	ascent := model.Coordinate(float64(fontDescriptor.Ascent) / 1000 * fontSize)
	descent := model.Coordinate(float64(fontDescriptor.Descent) / 1000 * fontSize)
	lenText := model.Coordinate(p.Pdf.GetStringWidth(t.Text))

	switch t.Orientation.Align {
	case "center":
		ancor.SetX(ancor.X() - lenText/2)
	case "right":
		ancor.SetX(ancor.X() - lenText)
	}

	tlPoint := model.NewPoint(ancor.X()-t.Orientation.Space.Left(), ancor.Y()-descent-ascent-t.Orientation.Space.Top())
	brPoint := model.NewPoint(ancor.X()+lenText+t.Orientation.Space.Right(), ancor.Y()-descent+t.Orientation.Space.Bottom())
	cell := model.NewCell(*tlPoint, *brPoint, 1, "", "Black")

	if strings.Contains(t.Orientation.Border, "l") {
		p.PrintLine(*cell.BLeft())
	}
	if strings.Contains(t.Orientation.Border, "t") {
		p.PrintLine(*cell.BTop())
	}
	if strings.Contains(t.Orientation.Border, "r") {
		p.PrintLine(*cell.BRight())
	}
	if strings.Contains(t.Orientation.Border, "b") {
		p.PrintLine(*cell.BBottom())
	}
	if strings.Contains(t.Orientation.Border, "a") {
		for _, l := range cell.Lines {
			p.PrintLine(l)
		}
	}

	switch {
	case strings.Contains(t.Line, "u"):
		p.PrintLine(*model.NewHorisontLine(t.Orientation.Start.Y()+model.Coordinate(fontSize)/10, ancor.X(), ancor.X()+lenText, model.Coordinate(fontSize/15)))
	case strings.Contains(t.Line, "c"):
		p.PrintLine(*model.NewHorisontLine(t.Orientation.Start.Y()+descent, ancor.X(), ancor.X()+lenText, model.Coordinate(fontSize/15)))
	case strings.Contains(t.Line, "o"):
		p.PrintLine(*model.NewHorisontLine(t.Orientation.Start.Y()-descent-ascent-model.Coordinate(fontSize)/10, ancor.X(), ancor.X()+lenText, model.Coordinate(fontSize/15)))
	}

	p.Pdf.Text(
		float64(ancor.X()),
		float64(ancor.Y()),
		t.Text,
	)

	return *ancor.SetY(ancor.Y() + model.Coordinate(fontSize*t.Orientation.Padding))
}
