// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodels

import (
	"fmt"
	"io"
	"strconv"
)

type MenuInput struct {
	ItemName      string  `json:"itemName"`
	Price         float64 `json:"price"`
	IsStopSelling bool    `json:"isStopSelling"`
}

type PriceUnit string

const (
	// 台幣
	PriceUnitNtd PriceUnit = "NTD"
	// 日幣
	PriceUnitJpy PriceUnit = "JPY"
)

var AllPriceUnit = []PriceUnit{
	PriceUnitNtd,
	PriceUnitJpy,
}

func (e PriceUnit) IsValid() bool {
	switch e {
	case PriceUnitNtd, PriceUnitJpy:
		return true
	}
	return false
}

func (e PriceUnit) String() string {
	return string(e)
}

func (e *PriceUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PriceUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PriceUnit", str)
	}
	return nil
}

func (e PriceUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
