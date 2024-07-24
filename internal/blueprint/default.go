package blueprint

import (
	"secretary/internal/blueprint/modal"
	"secretary/internal/printer/model"
)

type DefaultBluprint struct {
	NameDoc string
	model.Options
	modal.Content
}

func (bp DefaultBluprint) Use() modal.Bluprint {
	return bp
}

func NewDefaultBluprint() modal.Bluprint {
	bp := DefaultBluprint{
		NameDoc: "DefaultDoc",
		Content: map[string]string{
			"title":       "DefaultDoc Bluprint",
			"description": "this bluprint show all default fiches.",
			"longText":    "long-long-long long long long long long long long very long long long long long long long long long long very very very very most very long text. This text do not see in once screen.",
			"shortText":   "st"},
	}
	return bp
}
