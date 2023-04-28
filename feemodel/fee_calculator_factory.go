package feemodel

import (
	"sahaj-parking-lot/enum"
)

type FeeFactory interface {
	GetFeeCalculator(parkingType enum.ParkingLotType) FeeCalculator
}

type feeFactory struct {
	mallFeeCalculator    FeeCalculator
	stadiumFeeCalculator FeeCalculator
	airportFeeCalculator FeeCalculator
}

func NewFeeFactory() FeeFactory {
	return feeFactory{
		mallFeeCalculator:    NewMallFeeCalculator(),
		stadiumFeeCalculator: NewStadiumFeeCalculator(),
		airportFeeCalculator: NewAirportFeeCalculator(),
	}
}

func (f feeFactory) GetFeeCalculator(parkingType enum.ParkingLotType) FeeCalculator {
	switch parkingType {
	case enum.Mall:
		return f.mallFeeCalculator
	case enum.Airport:
		return f.airportFeeCalculator
	case enum.Stadium:
		return f.stadiumFeeCalculator
	}

	return nil
}
