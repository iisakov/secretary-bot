package pdfer

import "secretary/internal/printer/model"

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
