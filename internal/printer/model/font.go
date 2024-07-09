package model

import (
	"encoding/json"
	"os"
)

type Font struct {
	Family string `json:"family"`
	Style  string `json:"style"`
	File   string `json:"file"`
}

type Fonts []Font

func FontLoad() (*Fonts, error) {
	var fonts = Fonts{}

	plan, err := os.ReadFile("../../source/fonts/fonts_codes.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(plan, &fonts)
	if err != nil {
		return nil, err
	}

	return &fonts, nil
}

type FontVariant struct {
	Name   string
	Family string
	Style  string
	Size   float64
	Color  string
}

type FontVariants []FontVariant

func FontVariantsLoad() (*FontVariants, error) {
	var fv = FontVariants{}

	plan, err := os.ReadFile("../../source/fonts/font_variants_codes.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(plan, &fv)
	if err != nil {
		return nil, err
	}

	return &fv, nil
}

func (fvs FontVariants) FindByName(name string) (*FontVariant, bool) {
	for _, fv := range fvs {
		if fv.Name == name {
			return &fv, true
		}
	}
	return fvs.Default(), false
}

func (fvs FontVariants) Default() *FontVariant {
	fv, _ := fvs.FindByName("Default")
	return fv
}
