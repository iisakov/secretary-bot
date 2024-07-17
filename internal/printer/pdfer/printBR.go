package pdfer

import (
	"secretary/internal/printer/model"
	"strings"
)

// TODO Не работает правая граница
func (p PDFer) PrintTextBR(t model.Text, maxlen model.Coordinate) model.Point {
	p.SetFont(t.FVName)
	tBuilder := model.NewTextBuilder()

	textPart := ""
	startPoint := *model.NewPoint(t.Orientation.Start.X(), t.Orientation.Start.Y())
	indent := t.Orientation.Indent.Indent
	numLine := 0
	for _, ts := range strings.Split(t.Text, " ") {
		if p.Pdf.GetStringWidth(textPart) > float64(maxlen-indent)-p.Pdf.GetStringWidth(ts) {
			startPoint = printTextPart(p, t, tBuilder, startPoint, textPart, indent)
			textPart = ""
			numLine++
		}
		textPart += ts + " "
		switch {
		case numLine < int(t.Orientation.Indent.NumLines):
			switch t.Orientation.Align {
			case "right":
				indent = -t.Orientation.Indent.Indent
			default:
				indent = t.Orientation.Indent.Indent
			}
		default:
			indent = 0
		}
	}
	if textPart != "" {
		startPoint = printTextPart(p, t, tBuilder, startPoint, textPart, indent)
	}

	return startPoint
}

func printTextPart(
	p model.Printer,
	t model.Text,
	tb model.TextBuilderable,
	sp model.Point,
	tp string,
	i model.Coordinate) model.Point {
	sp.SetX(sp.X() + i)
	t.Orientation.Start = sp
	ct := tb.Name(t.Name).FVName(t.FVName).Orientation(t.Orientation).Text(tp).Line(t.Line).Build()
	sp = p.PrintText(ct)
	sp.SetX(sp.X() - i)
	return sp

}
