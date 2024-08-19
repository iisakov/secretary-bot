package blueprint

import (
	"encoding/json"
	"fmt"
	"secretary/internal/blueprint/model"
	bpm "secretary/internal/blueprint/model"
	pm "secretary/internal/printer/model"
	"strconv"
	"strings"
	"time"
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

	pointRedline := *pm.NewPoint(p.GetPageSize().X()/2, p.GetPageSize().Y()/10)

	tBuilder := pm.NewTextBuilder().FVName("fgisH1").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "center"}).Line("").Text("СВИДЕТЕЛЬСТВО О ПОВЕРКЕ СРЕДСТВА ИЗМЕРЕНИЙ")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	pointHorizontLine := pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 1.5, Align: "center"}).Line("br").Text(bp.Cert.Organization)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-4), Padding: 1, Align: "center"}).Line("").Text("наименование аккредитованного в соответствии с законодательством Российской Федерации об аккредитации в национальной системе аккредитации юридического лица или индивидуального предпринимателя, выполнившего поверку")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(10), Padding: 0, Align: "center"}).Line("").Text("Уникальный номер записи об аккредитации в реестре аккредитованных лиц          ")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	p.PrintLine(*pm.NewHorisontLine(pointRedline.Y()+2, p.GetPageSize().X()-150, p.GetPageSize().X()-100, 0.5))

	tBuilder = tBuilder.FVName("fgisH2").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(30), Padding: 2, Align: "center"}).Line("").Text("СВИДЕТЕЛЬСТВО О ПОВЕРКЕ №: " + bp.Cert.CertNum)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftX(100), Padding: 2.5, Align: "left"}).Line("").Text("Действительно до: " + bp.Cert.ValidDate)
	pointRedline = p.PrintText(tBuilder.Build())

	//Средство измерения
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50), Padding: 0, Align: "left"}).Line("").Text("Средство измерения:")
	pointRedline = p.PrintText(tBuilder.Build())

	pointDescription := pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(100), NumLines: 1}}).Line("br").Text(bp.Cert.MI)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	p.PrintLine(*pm.NewHorisontLine(pointHorizontLine.ShiftY(8*2+0.8).Y(), 50, p.GetPageSize().X()-50+1, 8.0/15.0))

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2).ShiftY(7), Padding: 2.7, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(70), NumLines: 1}}).Line("").Text("наименование и обозначение типа, модификация (при наличии) средства измерений, регистрационный номер в Федеральном информационном фонде по обеспечению единства измерений, присвоенный при утверждении типа")
	pointDescription = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//Заводской номер
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 0, Align: "left"}).Line("").Text("Заводской номер:")
	pointRedline = p.PrintText(tBuilder.Build())

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(100), NumLines: 1}}).Line("br").Text(bp.Cert.ManufactureNum)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2).ShiftY(7), Padding: 2.7, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(70), NumLines: 1}}).Line("").Text("заводской (серийный) номер или буквенно-цифровое обозначение")
	pointDescription = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//в составе поверено
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 0, Align: "left"}).Line("").Text("в составе поверено")
	pointRedline = p.PrintText(tBuilder.Build())

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(100), NumLines: 1}}).Line("br").Text(bp.Cert.BriefCharacteristics)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	p.PrintLine(*pm.NewHorisontLine(pointHorizontLine.ShiftY(8*2+0.8).Y(), 50, p.GetPageSize().X()-50+1, 8.0/15.0))

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2).ShiftY(7), Padding: 2.7, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(60), NumLines: 1}}).Line("").Text("наименование единиц величин, диапазонов измерений, на которых поверено средство измерений или которые исключены из поверки")
	pointDescription = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//в соответствии с
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 0, Align: "left"}).Line("").Text("в соответствии с")
	pointRedline = p.PrintText(tBuilder.Build())

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(100), NumLines: 1}}).Line("br").Text(bp.Cert.DocTitle)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2).ShiftY(7), Padding: 2.7, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(60), NumLines: 1}}).Line("").Text("наименование или обозначение документа, на основании которого выполнена поверка")
	pointDescription = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//с применением эталонов
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 0, Align: "left"}).Line("").Text("с применением эталонов")
	pointRedline = p.PrintText(tBuilder.Build())

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: pointRedline, Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(120), NumLines: 1}}).Line("br").Text(strings.Join(bp.Cert.Mieta, "; "))
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	p.PrintLine(*pm.NewHorisontLine(pointHorizontLine.ShiftY(8*2+0.8).Y(), 50, p.GetPageSize().X()-50+1, 8.0/15.0))

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointDescription.SetX(p.GetPageSize().X() / 2).ShiftY(7), Padding: 2.7, Align: "center", Indent: pm.Indent{Indent: pm.Coordinate(60), NumLines: 1}}).Line("").Text("регистрационные номера эталонов и (или) наименования и обозначения типов стандартных образцов и (или) средств измерений, заводские номера, обязательные требования к эталонам")
	pointDescription = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-120)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//при следующих значениях влияющих факторов:
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 1, Align: "left"}).Line("").Text("при следующих значениях влияющих факторов:")
	pointRedline = p.PrintTextBR(tBuilder.Build(), 150)

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-8), Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(120), NumLines: 1}}).Line("br").Text("указано в протоколе, соответствуют МП")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//Поверитель
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 1, Align: "left"}).Line("").Text("Поверитель:")
	pointRedline = p.PrintTextBR(tBuilder.Build(), 150)

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-8), Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(120), NumLines: 1}}).Line("br").Text("указано в протоколе")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//Постоянный адрес записи сведений о результатах поверки в ФИФ ОЕИ:
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 1, Align: "left"}).Line("").Text("Постоянный адрес записи сведений о результатах поверки в ФИФ ОЕИ:")
	pointRedline = p.PrintTextBR(tBuilder.Build(), 150)

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-8), Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(120), NumLines: 1}}).Line("br").Text(bp.Cert.Url + " ")
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//Номер записи сведений о результатах поверки в ФИФ ОЕИ:
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 1, Align: "left"}).Line("").Text("Номер записи сведений о результатах поверки в ФИФ ОЕИ:")
	pointRedline = p.PrintTextBR(tBuilder.Build(), 150)

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-8), Padding: 2, Align: "left", Indent: pm.Indent{Indent: pm.Coordinate(120), NumLines: 1}}).Line("br").Text(bp.Cert.RecordsNum)
	pointRedline = p.PrintTextBR(tBuilder.Build(), p.GetPageSize().X()-50)

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	//Дата поверки
	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.SetX(50).ShiftY(8), Padding: 1, Align: "left"}).Line("").Text("Дата поверки:")
	pointRedline = p.PrintTextBR(tBuilder.Build(), 150)

	pointDescription = pointRedline
	pointHorizontLine = pointRedline

	tBuilder = tBuilder.FVName("fgisBody").Orientation(pm.Orientation{Start: *pointRedline.ShiftY(-8).SetX(170), Padding: 2, Align: "left"}).Line("u").Text(bp.Cert.VrfDate)
	pointRedline = p.PrintText(tBuilder.Build())

	if pointDescription.Y() > pointRedline.Y() {
		pointRedline.SetY(pointDescription.Y())
	}

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointRedline.SetY(p.GetPageSize().Y() - 70).SetX(p.GetPageSize().X() / 4), Padding: 1, Align: "center"}).Line("").Text("должность руководителя или другого уполномоченного лица")
	pointHorizontLine = p.PrintTextBR(tBuilder.Build(), 200)

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointRedline.SetY(p.GetPageSize().Y() - 70).SetX(p.GetPageSize().X() / 2), Padding: 1, Align: "center"}).Line("").Text("подпись")
	p.PrintTextBR(tBuilder.Build(), 200)

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointRedline.SetY(p.GetPageSize().Y() - 70).SetX(p.GetPageSize().X() - p.GetPageSize().X()/4), Padding: 1, Align: "center"}).Line("").Text("фамилия, инициалы")
	p.PrintTextBR(tBuilder.Build(), 200)

	p.PrintLine(*pm.NewHorisontLine(pointHorizontLine.ShiftY(-20).Y(), 50, p.GetPageSize().X()-50+1, 8.0/15.0))

	tBuilder = tBuilder.FVName("fgisDescBody").Orientation(pm.Orientation{Start: *pointRedline.SetY(p.GetPageSize().Y() - 25).SetX(p.GetPageSize().X() / 2), Padding: 2, Align: "center"}).Line("").Text(fmt.Sprintf("Выписка о результатах поверки СИ № %s сформирована автоматически %s по данным, содержащимся в ФИФ ОЕИ", bp.Cert.CertNum, time.Now().Format("02.01.2006 03:04")))
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
			"Рег. №" + bp.EnrichedСontent.MiInfo.EtaMI.MitypeNumber}, ";")
	case bp.EnrichedСontent.MiInfo.SingleMI != model.SingleMI{}:
		mi = strings.Join([]string{
			bp.EnrichedСontent.MiInfo.SingleMI.MitypeTitle,
			bp.EnrichedСontent.MiInfo.SingleMI.MitypeType,
			bp.EnrichedСontent.MiInfo.SingleMI.Modification,
			"Рег. №" + bp.EnrichedСontent.MiInfo.SingleMI.MitypeNumber}, ";")
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
		RecordsNum:           c["result"].(map[string]any)["recordsNum"].(string),
		VrfDate:              bp.EnrichedСontent.VriInfo.VrfDate,
	}

	fmt.Println(c["result"].(map[string]any)["recordsNum"], bp.Cert.RecordsNum, bp.RecordsNum)

	return bp
}
