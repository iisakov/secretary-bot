package blueprint

import (
	bpm "secretary/internal/blueprint/model"
	pm "secretary/internal/printer/model"
)

type DefaultBluprint struct {
	NameDoc string
	pm.Options
	bpm.Content
}

func (bp DefaultBluprint) Use(p pm.Printer) bpm.Bluprint {
	p.SetColor("Pastel gray")
	p.PrintLine(*pm.NewVerticalLine(20, 0, p.GetPageSize().Y(), 1))
	p.PrintLine(*pm.NewVerticalLine(p.GetPageSize().X()-20, 0, p.GetPageSize().Y(), 1))

	pointRedline := *pm.NewPoint(p.GetPageSize().X()/2, 20)

	tBuilder := pm.NewTextBuilder().FVName("DH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 0.7, Align: "center"}).Line("u").Text(bp.Content["title"])
	pointRedline = p.PrintText(tBuilder.Build())

	tBuilder = tBuilder.FVName("DDescH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 1, Align: "center"}).Line("").Text(bp.Content["description"])
	pointRedline = p.PrintText(tBuilder.Build())

	// longText section
	pointRedline.SetX(20).ShiftY(11.0)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 1}).Line("").Text("LongText:")
	pointDescription := p.PrintText(tBuilder.Build())
	tBuilder = tBuilder.FVName("DDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2), Padding: 2.5, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(40), NumLines: 1}}).Line("").Text("it is intended to show how sections of a document can be signed")
	p.PrintTextBR(tBuilder.Build(), 200)

	pointRedline.ShiftX(15)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Indent: pm.Indent{Indent: pm.Coordinate(40), NumLines: 1}}).Line("br").Text(bp.Content["longText"] + " 01")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-20)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2}).Line("br").Text(bp.Content["longText"] + " 02")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-20)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2}).Line("br").Text(bp.Content["longText"] + " 03")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-20)

	p.OutputDoc(bp.NameDoc)
	return bp
}

func (bp DefaultBluprint) GetOptions() pm.Options {
	return bp.Options
}

func NewDefaultBluprint() bpm.Bluprint {
	bp := DefaultBluprint{
		NameDoc: "DefaultDoc",
		Options: pm.Options{
			Orientation: "P",
			Unit:        "pt",
			Size:        "A4",
			FontDir:     "",
		},
		Content: map[string]string{
			"title":       "DefaultDoc Bluprint",
			"description": "this bluprint show all default fiches.",
			"longText":    "long-long-long long long long long long long long very long long long long long long long long long long very very very very most very long text. This text do not see in once screen.",
			"shortText":   "st"},
	}
	return bp
}
