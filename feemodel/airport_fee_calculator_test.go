package feemodel

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestAirportFeeCalculator_Calculate_For_TwoWheelers(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportCalculator := testFeeFactory.GetFeeCalculator(enum.Airport)
	charges := airportCalculator.Calculate(time.Now().Add(-55*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(0))
	charges = airportCalculator.Calculate(
		time.Now().Add(-14*time.Hour).Add(-59*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(60))
	charges = airportCalculator.Calculate(time.Now().Add(-36*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(160))
	charges = airportCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(40))
}

func TestAirportFeeCalculator_Calculate_For_BigFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	mallCalculator := testFeeFactory.GetFeeCalculator(enum.Airport)
	charges := mallCalculator.Calculate(time.Now().Add(-10*time.Minute), enum.BigFourWheeler)
	assert.Equal(t, charges, float64(-1))
}

func TestAirportFeeCalculator_Calculate_For_SmallFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportParkingCalculator := testFeeFactory.GetFeeCalculator(enum.Airport)
	charges := airportParkingCalculator.Calculate(
		time.Now().Add(-50*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(60))
	charges = airportParkingCalculator.Calculate(
		time.Now().Add(-23*time.Hour).Add(-59*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(80))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-73*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(400))
}
