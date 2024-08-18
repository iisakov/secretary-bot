package model

import pm "secretary/internal/printer/model"

type Bluprint interface {
	Use(pm.Printer) Bluprint
	GetOptions() pm.Options
	SetContent(c Content) Bluprint
}

type Content map[string]any
