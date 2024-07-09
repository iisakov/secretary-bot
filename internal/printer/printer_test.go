package printer

import (
	"fmt"
	"secretary/internal/printer/model"
	"secretary/internal/printer/pdfer"
	"testing"
)

func TestPDFPrintText(t *testing.T) {
	p := NewPrinter(
		pdfer.PDFer{},
		"PDFerTest",
		model.Options{
			Orientation: "P",
			Unit:        "pt",
			Size:        "A4",
			FontDir:     "",
			Inks:        []model.Ink{{Name: "test", Color: [3]int{100, 111, 255}}},
			Fonts:       []model.Font{{Family: "PT-Root-UI", Style: "", File: "../../source/fonts/PT/PT-Root-UI/pt-root-ui_regular.ttf"}},
		},
	)

	y := model.Coordinate(10)

	for i, fv := range *p.GetFontVariants() {
		line := []string{"u", "c", "o"}[i%3]
		// Текст с выравниванием по левому краю
		y = print(p, line, "", fv, y, 10, "left", model.Space{})
		// Текст с выравниванием по центру
		y = print(p, line, "", fv, y, p.GetPageSize().X()/2, "center", model.Space{})
		// Текст с выравниванием по правому краю
		y = print(p, line, "", fv, y, p.GetPageSize().X()-10, "right", model.Space{})
	}

	for i, fv := range *p.GetFontVariants() {
		// Текст с границами
		space := model.Space{5 * model.Coordinate(i), model.Coordinate(fv.Size), 5 * model.Coordinate(i), model.Coordinate(fv.Size)}
		y = print(p, "", "tl", fv, y, 10, "left", space)
		y += model.Coordinate(fv.Size) * 2
		y = print(p, "", "lr", fv, y, p.GetPageSize().X()/2, "center", space)
		y += model.Coordinate(fv.Size) * 2
		y = print(p, "", "a", fv, y, p.GetPageSize().X()-10, "right", space)
		y += model.Coordinate(fv.Size) * 2
	}

	if err := p.OutputDoc("PDFerTest"); err != nil {
		fmt.Println(err)
	}
}

func print(
	p model.Printer,
	line, border string,
	fv model.FontVariant,
	y, x model.Coordinate,
	align string,
	space model.Space) model.Coordinate {

	if y >= p.GetPageSize().Y() {
		p.AddPage()
		y = 10
	}

	txt := model.Text{
		Name:   "TEST",
		FVName: fv.Name,
		Orientation: model.Orientation{
			Start:  *model.NewPoint(x, 10+y),
			Align:  align,
			Border: border,
			Space:  space},
		Text: fmt.Sprintf("%s %s: %s %s c: %s", "`,y", fv.Name, fv.Family, fv.Style, fv.Color),
		Line: line}

	p.PrintText(txt)

	y += model.Coordinate(fv.Size)
	return y
}
