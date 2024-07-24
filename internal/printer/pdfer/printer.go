package pdfer

import (
	"fmt"
	"log"
	"os"
	"secretary/internal/printer/model"

	"github.com/go-pdf/fpdf"
)

type PDFer struct {
	NameDoc      string
	Pdf          *fpdf.Fpdf
	Inks         model.Inks
	FontVariants model.FontVariants
}

func NewPDFer() model.Printer {
	return PDFer{}
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

func (p PDFer) PrintLine(l model.Line) model.Point {
	p.Pdf.SetLineWidth(float64(l.Width))
	p.Pdf.Line(float64(l.Start.X()), float64(l.Start.Y()), float64(l.Finish.X()), float64(l.Finish.Y()))

	return l.Start
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
