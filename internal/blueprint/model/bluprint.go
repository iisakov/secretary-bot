package model

import pm "secretary/internal/printer/model"

type Bluprint interface {
	Use(pm.Printer) Bluprint
	GetOptions() pm.Options
}

type Content map[string]string
