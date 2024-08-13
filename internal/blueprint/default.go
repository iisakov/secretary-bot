package blueprint

import (
	bpm "secretary/internal/blueprint/model"
	pm "secretary/internal/printer/model"
	"time"
)

type DefaultBluprint struct {
	NameDoc string
	pm.Options
	bpm.Content
}

func (bp DefaultBluprint) Use(p pm.Printer) bpm.Bluprint {
	defer p.OutputDoc(bp.NameDoc)

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

	//shortText section
	pointRedline.SetX(20).ShiftY(20)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 1}).Line("").Text("ShortText:")
	pointDescription = p.PrintText(tBuilder.Build())
	tBuilder = tBuilder.FVName("DDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2), Padding: 2.5, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(40), NumLines: 1}}).Text("it is intended to show how sections of a document can be signed")
	p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-40)

	pointRedline.ShiftX(15)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Indent: pm.Indent{Indent: pm.Coordinate(40), NumLines: 1}}).Line("br").Text(bp.Content["shortText"] + " 01")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-20)

	pointRedline.ShiftY(20)
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextCenter"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "a", "Default").AddText(tBuilder.Build(), "center"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextTop"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "center top"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextLeft"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "left"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextRight"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "right"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextBottom"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "center bottom"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextTopLeft"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "top left"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextTopRight"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "top right"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextBottomLeft"] + " 01")
	p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "bottom left"))
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{}).Text(bp.Content["cellTextBottomRight"] + " 01")
	pointRedline = p.PrintCell(*pm.NewCell(*pointRedline.SetX(20), *pm.NewPoint(p.GetPageSize().X()-20, pointRedline.Y()+200), 1, "", "Default").AddText(tBuilder.Build(), "bottom right"))

	tBuilder = tBuilder.FVName("DH3").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(80).ShiftX(100), Padding: 0}).Line("").Text("Create time:")
	pointRedline = p.PrintText(tBuilder.Build())
	tBuilder = tBuilder.FVName("DBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftX(70), Padding: 0}).Line("u").Text(bp.Content["date"])
	pointRedline = p.PrintText(tBuilder.Build())

	tBuilder = tBuilder.FVName("DH3").Orientation(pm.Orientation{Start: *pointRedline.SetX(p.GetPageSize().X() - 200), Padding: 0, Align: "right"}).Line("").Text("Client signe:")
	pointRedline = p.PrintText(tBuilder.Build())
	p.PrintLine(*pm.NewHorisontLine(pointRedline.Y(), pointRedline.X()+5, pointRedline.X()+80, 1))

	tBuilder = tBuilder.FVName("DSigne").Orientation(pm.Orientation{Start: *pm.NewPoint(p.GetPageSize().X()-20, p.GetPageSize().Y()-20), Padding: 1, Align: "right"}).Text(bp.Content["signe"])
	pointRedline = p.PrintText(tBuilder.Build())
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
			"title":               "DefaultDoc Bluprint",
			"description":         "this bluprint show all default fiches.",
			"longText":            "long-long-long long long long long long long long very long long long long long long long long long long very very very very most very long text. This text do not see in once screen.",
			"shortText":           "st",
			"cellTextCenter":      "cell text center.",
			"cellTextLeft":        "cell text left.",
			"cellTextRight":       "cell text right.",
			"cellTextTop":         "cell text center top.",
			"cellTextBottom":      "cell text bottom.",
			"cellTextTopLeft":     "cell text top left",
			"cellTextTopRight":    "cell text top right",
			"cellTextBottomLeft":  "cell text bottom left",
			"cellTextBottomRight": "cell text bottom right",
			"date":                time.Now().Format("02.01.2006 15:04:05"),
			"signe":               "secretary bot [by_Artisan] v.2024.07.31:0.9.9"},
	}
	return bp
}
