package main

import (
	"fmt"
	"secretary/internal/blueprint"
	"secretary/internal/printer"
	"secretary/internal/printer/model"
)

func main() {
	_, _ = model.InksLoad()
	p := printer.NewDefaultPrinter(model.Options{})
	b := blueprint.NewDefaultBluprint()
	fmt.Println(p, b)

}
