package operator

import (
	bpm "secretary/internal/blueprint/model"
	om "secretary/internal/operator/model"
	pm "secretary/internal/printer/model"
	request "secretary/internal/requestor"
)

type FgisOperator struct {
	printer   pm.Printer
	blueprint bpm.Bluprint
	content   map[string]any
}

func NewFgisOperator(p pm.Printer, bp bpm.Bluprint) om.Operator {
	return FgisOperator{printer: p, blueprint: bp}
}

func (o FgisOperator) SetContent(s string) om.Operator {
	o.RequestContent(s)
	o.blueprint = o.blueprint.SetContent(o.content)
	return o
}

func (o FgisOperator) UseBluprint() {
	o.blueprint.Use(o.printer)
}

func (o *FgisOperator) RequestContent(certName string) (err error) {
	r, err := request.FgisRequest(certName)
	if err != nil {
		return err
	}
	o.printer = o.printer.Configurate(certName, o.blueprint.GetOptions())

	o.content = r

	return nil
}
