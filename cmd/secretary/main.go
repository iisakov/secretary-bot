package main

import (
	"secretary/internal/blueprint"
	"secretary/internal/operator"
	"secretary/internal/printer"
)

func main() {
	bp := blueprint.NewDefaultBluprint()
	p := printer.NewDefaultPrinter(bp.GetOptions())
	o := operator.NewDefaultOperator(p, bp)
	o.UseBluprint()
}
