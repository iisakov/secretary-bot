package printer

import (
	"secretary/internal/printer/model"
)

func NewPrinter(p model.Printer, nd string, o model.Options) model.Printer {
	return p.Configurate(nd, o)
}
