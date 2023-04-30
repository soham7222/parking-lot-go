package feemodel

import (
	"math"
	"sahaj-parking-lot/enum"
	"time"
)

type airportFeeCalculator struct {
}

func NewAirportFeeCalculator() FeeCalculator {
	return airportFeeCalculator{}
}

func (a airportFeeCalculator) Calculate(vehicleEntry time.Time, spotType enum.SpotType) float64 {
	timeSpentInHours := time.Now().Sub(vehicleEntry)
	switch spotType {
	case enum.TwoWheelers:
		return a.calculateForTwoWheelers(timeSpentInHours)
	case enum.SmallFourWheeler:
		return a.calculateForSmallFourWheelers(timeSpentInHours)
	default:
		return -1
	}
}

func (a airportFeeCalculator) calculateForTwoWheelers(timeSpent time.Duration) float64 {
	spentInMinutes := timeSpent.Minutes()
	timeSpentInDays := math.Ceil(timeSpent.Hours() / 24)
	if spentInMinutes >= 24*60 {
		return timeSpentInDays * 80
	} else if spentInMinutes > 8*60 {
		return 60
	} else if spentInMinutes > 1*60 {
		return 40
	} else {
		return 0
	}
}

func (a airportFeeCalculator) calculateForSmallFourWheelers(timeSpent time.Duration) float64 {
	spentInMinutes := timeSpent.Minutes()
	timeSpentInDays := math.Ceil(timeSpent.Hours() / 24)

	if spentInMinutes > 24*60 {
		return timeSpentInDays * 100
	} else if spentInMinutes > 12*60 {
		return 80
	} else {
		return 60
	}
}
