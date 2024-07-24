package operator

import (
	bpm "secretary/internal/blueprint/model"
	om "secretary/internal/operator/model"
	pm "secretary/internal/printer/model"
)

type DefaultOperator struct {
	printer   pm.Printer
	blueprint bpm.Bluprint
}

func NewDefaultOperator(p pm.Printer, bp bpm.Bluprint) om.Operator {
	return DefaultOperator{printer: p, blueprint: bp}
}

func (o DefaultOperator) UseBluprint() {
	o.blueprint.Use(o.printer)
}
