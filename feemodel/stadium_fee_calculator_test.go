package feemodel

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestStadiumFeeCalculator_Calculate_For_SmallFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportParkingCalculator := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := airportParkingCalculator.Calculate(time.Now().Add(-3*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(60))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(180))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-9*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(1080))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-15*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(3000))
}

func TestStadiumFeeCalculator_Calculate_For_TwoWheelers(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportParkingCalculator := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := airportParkingCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(90))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-9*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(540))
	charges = airportParkingCalculator.Calculate(time.Now().Add(-15*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(1500))
}

func TestStadiumFeeCalculator_Calculate_For_BigFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportParkingCalculator := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := airportParkingCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.BigFourWheeler)
	assert.Equal(t, charges, float64(-1))
}
