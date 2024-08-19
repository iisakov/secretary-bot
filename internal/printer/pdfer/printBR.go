package pdfer

import (
	"fmt"
	"secretary/internal/printer/model"
	"strings"
)

func (p PDFer) PrintTextBR(t model.Text, maxlen model.Coordinate) model.Point {
	p.SetFont(t.FVName)
	tBuilder := model.NewTextBuilder()
	fontSize, _ := p.Pdf.GetFontSize()
	avgTextPart := p.Pdf.GetStringWidth(t.Text) / float64(len(strings.Split(t.Text, " ")))

	textPart := ""
	startPoint := *model.NewPoint(t.Orientation.Start.X(), t.Orientation.Start.Y())
	indent := t.Orientation.Indent.Indent
	numLine := 0

	for _, ts := range strings.Split(t.Text, " ") {
		if p.Pdf.GetStringWidth(textPart) > float64(maxlen-indent)-avgTextPart*2.5 {
			switch t.Orientation.Align {
			case "right":
				if t.Line == "br" {
					p.PrintLine(*model.NewHorisontLine(startPoint.Y()+model.Coordinate(fontSize)/10, startPoint.X()+indent-model.Coordinate(fontSize)/7, maxlen+model.Coordinate(fontSize)/7, model.Coordinate(fontSize/15)))
				}
				startPoint = printTextPart(p, t, tBuilder, startPoint, textPart, -indent)
			default:
				if t.Line == "br" {
					if t.Orientation.Align == "center" {
						fmt.Println("center")
						sp := startPoint.X() + indent - model.Coordinate(fontSize)/7
						ep := maxlen + model.Coordinate(fontSize)/7
						p.PrintLine(*model.NewHorisontLine(startPoint.Y()+model.Coordinate(fontSize)/10, sp+sp-ep-20, ep+20, model.Coordinate(fontSize/15)))
					} else {
						p.PrintLine(*model.NewHorisontLine(startPoint.Y()+model.Coordinate(fontSize)/10, startPoint.X()+indent-model.Coordinate(fontSize)/7, maxlen+model.Coordinate(fontSize)/7, model.Coordinate(fontSize/15)))
					}
				}
				startPoint = printTextPart(p, t, tBuilder, startPoint, textPart, indent)
			}

			textPart = ""
			numLine++
		}
		textPart += ts + " "
		switch {
		case numLine < int(t.Orientation.Indent.NumLines):
			indent = t.Orientation.Indent.Indent
		default:
			indent = 0
		}
	}
	if textPart != "" {
		if t.Line == "br" {
			if t.Orientation.Align == "center" {
				fmt.Println("center")
				sp := startPoint.X() + indent - model.Coordinate(fontSize)/7
				ep := maxlen + model.Coordinate(fontSize)/7
				p.PrintLine(*model.NewHorisontLine(startPoint.Y()+model.Coordinate(fontSize)/10, sp+sp-ep-10, ep+10, model.Coordinate(fontSize/15)))
			} else {
				p.PrintLine(*model.NewHorisontLine(startPoint.Y()+model.Coordinate(fontSize)/10, startPoint.X()+indent-model.Coordinate(fontSize)/7, maxlen+model.Coordinate(fontSize)/7, model.Coordinate(fontSize/15)))
			}
		}
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
