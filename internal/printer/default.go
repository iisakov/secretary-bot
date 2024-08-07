package printer

import (
	"secretary/internal/printer/model"
	"secretary/internal/printer/pdfer"
)

func NewPrinter(p model.Printer, nd string, o model.Options) model.Printer {
	return p.Configurate(nd, o)
}

func NewDefaultPrinter(o model.Options) model.Printer {
	p := pdfer.NewPDFer()
	p = p.Configurate("Default", o)
	return p
}
