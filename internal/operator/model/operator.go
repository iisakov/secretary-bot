package model

type Operator interface {
	UseBluprint()
	SetContent(s string) Operator
}
