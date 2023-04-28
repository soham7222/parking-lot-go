package feemodel

import (
	"sahaj-parking-lot/enum"
	"time"
)

type FeeCalculator interface {
	Calculate(vehicleEntry time.Time, spotType enum.SpotType) float64
}
