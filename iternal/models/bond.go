package models

import (
	"time"
)

type Bond struct {
	Name           string
	IntrestRate    float64
	TrueIntrstRate float64
	Coupon         float64
	MarketValue    float64
	NominalValue   float64
	Aci            float64
	NearestPayDay  time.Time
}

func New(
	name string,
	intrestRate float64,
	coupon float64,
	marketValue float64,
	nominalValue float64,
	aci float64,
	nearestPayDay string) *Bond {
	npd, err := time.Parse("2006-01-02", nearestPayDay)
	if err != nil {
		return new(Bond)
	}
	return &Bond{
		Name:           name,
		IntrestRate:    intrestRate,
		TrueIntrstRate: coupon / marketValue,
		Coupon:         coupon,
		MarketValue:    marketValue,
		NominalValue:   nominalValue,
		Aci:            aci,
		NearestPayDay:  npd,
	}
}
