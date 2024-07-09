package pdfer

import (
	"fmt"
	"log"
	"os"
	"secretary/internal/printer/model"
	"strings"

	"github.com/go-pdf/fpdf"
)

type PDFer struct {
	NameDoc      string
	Pdf          *fpdf.Fpdf
	Inks         model.Inks
	FontVariants model.FontVariants
}

func (p PDFer) Configurate(nd string, o model.Options) model.Printer {
	inks, err := model.InksLoad()
	if err != nil {
		log.Fatal(err)
	}
	fonts, err := model.FontLoad()
	if err != nil {
		log.Fatal(err)
	}
	fontVariants, err := model.FontVariantsLoad()
	if err != nil {
		log.Fatal(err)
	}

	result := PDFer{
		NameDoc:      nd,
		Pdf:          fpdf.New(o.Orientation, o.Unit, o.Size, o.FontDir),
		Inks:         append(*inks, o.Inks...),
		FontVariants: append(*fontVariants, o.FontVariants...),
	}

	for _, f := range append(*fonts, o.Fonts...) {
		result.Pdf.AddUTF8Font(f.Family, f.Style, f.File)
	}

	result.Pdf.AddPage()

	return result
}

func (p PDFer) GetFontVariants() *model.FontVariants {
	return &p.FontVariants
}

func (p PDFer) GetPageSize() *model.Point {
	x, y := p.Pdf.GetPageSize()
	return model.NewPoint(model.Coordinate(x), model.Coordinate(y))
}

func (p PDFer) AddPage() {
	p.Pdf.AddPage()
}

func (p PDFer) SetColor(c string) {
	ink, ok := p.Inks.FindByColor(c)
	if !ok {
		ink = p.Inks.Black()
	}
	p.Pdf.SetTextColor(ink.Color[0], ink.Color[1], ink.Color[2])
	p.Pdf.SetDrawColor(ink.Color[0], ink.Color[1], ink.Color[2])
}

func (p PDFer) SetFont(n string) {
	fv, ok := p.FontVariants.FindByName(n)
	if !ok {
		fv = p.FontVariants.Default()
	}
	p.SetColor(fv.Color)
	p.Pdf.SetFont(fv.Family, fv.Style, fv.Size)
}

func (p PDFer) PrintLine(l model.Line) {
	p.Pdf.SetLineWidth(float64(l.Width))
	p.Pdf.Line(float64(l.Start.X()), float64(l.Start.Y()), float64(l.Finish.X()), float64(l.Finish.Y()))
}

func (p PDFer) PrintCell(cell model.Cell) {
	p.SetFont(cell.Text.FVName)

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

	fontDescriptor := p.GetFontDesc(cell.Text.FVName)
	fontSize, _ := p.Pdf.GetFontSize()
	ascent := model.Coordinate(float64(fontDescriptor.Ascent) / 1000 * fontSize)
	// descent := model.Coordinate(float64(fontDescriptor.Descent) / 1000 * fontSize)
	// lenText := model.Coordinate(p.Pdf.GetStringWidth(t.Text))
	ancor := model.NewPoint(cell.TL().X(), cell.TL().Y()+ascent)

	cell.Text.Orientation.Start = *ancor

	p.PrintText(cell.Text)
}

func (p PDFer) PrintText(t model.Text) {
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
}

func (p PDFer) PrintGrid(x, y model.Coordinate, opt string) {
	w, h := p.Pdf.GetPageSize()
	num := p.Pdf.PageCount()
	p.Pdf.PageNo()
	pageSize := model.NewPoint(model.Coordinate(w), model.Coordinate(h*float64(num)))

	switch opt {
	case "abs":
		for i := model.Coordinate(0); i < pageSize.Y(); i += y {
			l := model.NewHorisontLine(i, 0, model.Coordinate(w), 1)
			p.PrintLine(*l)
		}
		for i := model.Coordinate(0); i < pageSize.X(); i += x {
			l := model.NewVerticalLine(i, 0, model.Coordinate(h), 1)
			p.PrintLine(*l)
		}
	case "rel":
		for i := model.Coordinate(0); i < pageSize.Y(); i += pageSize.Y() / y {
			l := model.NewHorisontLine(i, 0, model.Coordinate(w), 1)
			p.PrintLine(*l)
		}
		for i := model.Coordinate(0); i < pageSize.X(); i += pageSize.X() / x {
			l := model.NewVerticalLine(i, 0, model.Coordinate(h), 1)
			p.PrintLine(*l)
		}
	}
}

func (p PDFer) GetFontDesc(FVName string) fpdf.FontDescType {
	fv, ok := p.GetFontVariants().FindByName(FVName)
	if !ok {
		p.GetFontVariants().Default()
	}

	return p.Pdf.GetFontDesc(fv.Family, fv.Style)
}

func (p PDFer) OutputDoc(domen string) error {
	_, err := os.Stat(fmt.Sprintf("../../out/%s", domen))
	if err != nil {
		os.Mkdir(fmt.Sprintf("../../out/%s", domen), 0777)
	}
	if err := p.Pdf.OutputFileAndClose(fmt.Sprintf("../../out/%s/%s.pdf", domen, p.NameDoc)); err != nil {
		return err
	}
	return nil
}
