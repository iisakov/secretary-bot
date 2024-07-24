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
	pointRedline := *pm.NewPoint(p.GetPageSize().X()/2, 20)

	tBuilder := pm.NewTextBuilder().FVName("DH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 1, Align: "center"}).Text(bp.Content["title"])
	pointRedline = p.PrintText(tBuilder.Build())

	tBuilder = tBuilder.FVName("DDesc").Orientation(pm.Orientation{Start: pointRedline, Padding: 1, Align: "center"}).Line("o").Text(bp.Content["description"])
	p.PrintText(tBuilder.Build())

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
