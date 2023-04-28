package feemodel

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestAirportFeeCalculator_Calculate_For_TwoWheelers(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	mallCalculator := testFeeFactory.GetFeeCalculator(enum.Airport)
	charges := mallCalculator.Calculate(time.Now().Add(-10*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(0))
	charges = mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(80))
	charges = mallCalculator.Calculate(time.Now().Add(-9*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(540))
	charges = mallCalculator.Calculate(time.Now().Add(-72*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(5760))
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
	charges := airportParkingCalculator.Calculate(time.Now().Add(-9*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(540))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-18*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(1440))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-30*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(3000))
}
