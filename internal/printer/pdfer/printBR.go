package pdfer

import "secretary/internal/printer/model"

func (p PDFer) PrintTextBR(t model.Text) model.Point {

	return *model.NewPoint(1, 1)
}
