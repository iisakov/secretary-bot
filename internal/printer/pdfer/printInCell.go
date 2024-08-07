package pdfer

import (
	"secretary/internal/printer/model"
	"strings"
)

func (p PDFer) PrintCell(cell model.Cell) model.Point {
	p.SetFont(cell.Text.FVName)
	p.SetColor(cell.Color)

	fontDescriptor := p.GetFontDesc(cell.Text.FVName)
	fontSize, _ := p.Pdf.GetFontSize()
	ascent := model.Coordinate(float64(fontDescriptor.Ascent) / 1000 * fontSize)
	descent := model.Coordinate(float64(fontDescriptor.Descent) / 1000 * fontSize)
	// lenText := model.Coordinate(p.Pdf.GetStringWidth(cell.Text.Text))
	ancor := model.NewPoint(cell.TL().X()+2, cell.TL().Y()+cell.Hight()/2-descent)

	if strings.Contains(cell.TextAlign, "left") {
		ancor.SetX(ancor.X() + cell.Text.Orientation.Space.Left())
	}
	if strings.Contains(cell.TextAlign, "center") {
		ancor.SetX(ancor.X() + cell.Width()/2)
		cell.Text.Orientation.Align = "center"
	}
	if strings.Contains(cell.TextAlign, "right") {
		ancor.SetX(ancor.X() + cell.BTop().Len() - 4 - cell.Text.Orientation.Space.Right())
		cell.Text.Orientation.Align = "right"
	}
	if strings.Contains(cell.TextAlign, "top") {
		ancor.SetY(ancor.Y() - cell.Hight()/2 + descent + ascent)
	}
	if strings.Contains(cell.TextAlign, "bottom") {
		ancor.SetY(ancor.Y() + cell.Hight()/2 + 2*descent)
	}

	cell.Text.Orientation.Start = *ancor

	p.PrintText(cell.Text)

	if strings.Contains(cell.Border, "l") {
		p.PrintLine(*cell.BLeft())
	}
	if strings.Contains(cell.Border, "t") {
		p.PrintLine(*cell.BTop())
	}
	if strings.Contains(cell.Border, "r") {
		p.PrintLine(*cell.BRight())
	}
	if strings.Contains(cell.Border, "b") {
		p.PrintLine(*cell.BBottom())
	}
	if strings.Contains(cell.Border, "a") {
		for _, l := range cell.Lines {
			p.PrintLine(l)
		}
	}

	return *cell.BL().ShiftY(float64(cell.Text.Orientation.Padding))
}
