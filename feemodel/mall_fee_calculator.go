package feemodel

import (
	"math"
	"sahaj-parking-lot/enum"
	"time"
)

type mallFeeCalculator struct {
}

func NewMallFeeCalculator() FeeCalculator {
	return mallFeeCalculator{}
}

func (m mallFeeCalculator) Calculate(vehicleEntry time.Time, spotType enum.SpotType) float64 {
	timeSpentInHours := math.Round(time.Now().Sub(vehicleEntry).Hours())
	timeSpentInMinutes := math.Round(time.Now().Sub(vehicleEntry).Minutes())
	if timeSpentInHours == 0 && timeSpentInMinutes != 0 {
		timeSpentInHours = 1
	}

	switch spotType {
	case enum.TwoWheelers:
		return 10 * timeSpentInHours
	case enum.SmallFourWheeler:
		return 20 * timeSpentInHours
	case enum.BigFourWheeler:
		return 50 * timeSpentInHours
	default:
		return -1
	}
}
