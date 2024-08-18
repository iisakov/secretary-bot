package blueprint

import (
	"encoding/json"
	"secretary/internal/blueprint/model"
	bpm "secretary/internal/blueprint/model"
	pm "secretary/internal/printer/model"
	"strconv"
	"strings"
)

type FgisBluprint struct {
	NameDoc string
	pm.Options
	bpm.Content
	EnrichedСontent model.Result
	model.Cert
}

func NewFgisBluprint() bpm.Bluprint {
	bp := FgisBluprint{
		NameDoc: "fgisDoc",
		Options: pm.Options{
			Orientation: "P",
			Unit:        "pt",
			Size:        "A4",
			FontDir:     "",
		},
	}
	return bp
}

func (bp FgisBluprint) Use(p pm.Printer) bpm.Bluprint {
	defer p.OutputDoc(bp.NameDoc)

	pointRedline := *pm.NewPoint(p.GetPageSize().X()/2, 20)

	tBuilder := pm.NewTextBuilder().FVName("DH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 0.7, Align: "center"}).Line("u").Text(bp.Cert.DocTitle)
	pointRedline = p.PrintText(tBuilder.Build())

	tBuilder = tBuilder.FVName("DDescH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 1, Align: "center"}).Line("").Text(bp.CertNum)
	pointRedline = p.PrintText(tBuilder.Build())

	return nil
}

func (bp FgisBluprint) GetOptions() pm.Options {
	return bp.Options
}

func (bp FgisBluprint) SetContent(c bpm.Content) bpm.Bluprint {
	bp.Content = c

	jsonString, err := json.Marshal(c["result"])

	if err = json.Unmarshal(jsonString, &bp.EnrichedСontent); err != nil {
		return nil
	}

	var MietaSlice []string
	for _, mieta := range bp.EnrichedСontent.Means.Mieta {
		MietaSlice = append(
			MietaSlice,
			strings.Join([]string{
				mieta.MitypeNumber,
				mieta.MitypeTitle,
				mieta.Notation,
				mieta.ManufactureNum,
				strconv.Itoa(mieta.ManufactureYear),
				mieta.RankTitle,
				mieta.SchemaTitle}, " "))
	}

	var bc, mi string
	if bp.EnrichedСontent.Info.BriefIndicator {
		bc = bp.EnrichedСontent.Info.BriefCharacteristics
	} else {
		bc = "В полном объёме"
	}

	switch {
	case bp.EnrichedСontent.MiInfo.EtaMI != model.EtaMI{}:
		mi = strings.Join([]string{
			bp.EnrichedСontent.MiInfo.EtaMI.MitypeTitle,
			bp.EnrichedСontent.MiInfo.EtaMI.MitypeType,
			bp.EnrichedСontent.MiInfo.EtaMI.Modification,
			"Рег. № " + bp.EnrichedСontent.MiInfo.EtaMI.MitypeNumber}, ";")
	case bp.EnrichedСontent.MiInfo.SingleMI != model.SingleMI{}:
		mi = strings.Join([]string{
			bp.EnrichedСontent.MiInfo.SingleMI.MitypeTitle,
			bp.EnrichedСontent.MiInfo.SingleMI.MitypeType,
			bp.EnrichedСontent.MiInfo.SingleMI.Modification,
			"Рег. № " + bp.EnrichedСontent.MiInfo.SingleMI.MitypeNumber}, ";")
	}

	bp.Cert = model.Cert{
		Organization:         bp.EnrichedСontent.VriInfo.Organization,
		ValidDate:            bp.EnrichedСontent.VriInfo.ValidDate,
		CertNum:              bp.EnrichedСontent.VriInfo.Applicable.CertNum,
		MI:                   mi,
		ManufactureNum:       bp.EnrichedСontent.MiInfo.EtaMI.ManufactureNum,
		DocTitle:             bp.EnrichedСontent.VriInfo.DocTitle,
		Mieta:                MietaSlice,
		Url:                  bp.EnrichedСontent.Url,
		BriefCharacteristics: bc,
	}

	return bp
}
