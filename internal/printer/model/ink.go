package model

import (
	"encoding/json"
	"os"
)

type Ink struct {
	Name  string `json:"label"`
	Color [3]int `json:"rgb"`
}

type Inks []Ink

func (inks Inks) FindByColor(color string) (*Ink, bool) {
	for _, i := range inks {
		if i.Name == color {
			return &i, true
		}
	}
	return nil, false
}

func (inks Inks) Black() *Ink {
	ink, _ := inks.FindByColor("Black")
	return ink
}

func InksLoad() (*Inks, error) {
	var inks = Inks{}

	plan, err := os.ReadFile("../../source/colors/color_codes.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(plan, &inks)
	if err != nil {
		return nil, err
	}

	return &inks, nil
}
