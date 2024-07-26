package models

import (
	"time"

)

type Bond struct {
	Name string
	IntrestRate float64
	TrueIntrstRate float64
	Cupon float64
	MarketValue float64
	NominalValue float64
	Aci float64
	NearestPayDay time.Time
}


