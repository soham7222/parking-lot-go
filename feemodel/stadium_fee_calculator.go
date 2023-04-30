package feemodel

import (
	"math"
	"sahaj-parking-lot/enum"
	"time"
)

type stadiumFeeCalculator struct {
}

func NewStadiumFeeCalculator() FeeCalculator {
	return stadiumFeeCalculator{}
}

func (s stadiumFeeCalculator) Calculate(vehicleEntry time.Time, spotType enum.SpotType) float64 {
	timeSpent := time.Now().Sub(vehicleEntry)
	switch spotType {
	case enum.TwoWheelers:
		return s.calculate(timeSpent,
			30,
			60,
			100)
	case enum.SmallFourWheeler:
		return s.calculate(timeSpent,
			60,
			120,
			200)
	default:
		return -1
	}
}

func (s stadiumFeeCalculator) calculate(timeSpent time.Duration,
	firstIntervalFee float64,
	secondIntervalFee float64,
	thirdIntervalFee float64) float64 {

	minutesSpent := timeSpent.Minutes()
	hoursSpent := math.Ceil(timeSpent.Hours())

	if minutesSpent > 12*60 {
		return firstIntervalFee + secondIntervalFee + (hoursSpent-12)*thirdIntervalFee
	} else if minutesSpent > 4*60 {
		return firstIntervalFee + secondIntervalFee
	} else {
		return firstIntervalFee
	}
}
