package feemodel

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestStadiumFeeCalculator_Calculate_For_SmallFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	stadiumCalculator := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := stadiumCalculator.Calculate(time.Now().Add(-3*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(60))
	charges = stadiumCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(60))
	charges = stadiumCalculator.Calculate(
		time.Now().Add(-11*time.Hour).Add(-30*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(180))
	charges = stadiumCalculator.Calculate(
		time.Now().Add(-13*time.Hour).Add(-5*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(580))
}

func TestStadiumFeeCalculator_Calculate_For_TwoWheelers(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	stadiumCalculation := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := stadiumCalculation.Calculate(
		time.Now().Add(-3*time.Hour).Add(-40*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(30))
	charges = stadiumCalculation.Calculate(
		time.Now().Add(-9*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(90))
	charges = stadiumCalculation.Calculate(
		time.Now().Add(-14*time.Hour).Add(-59*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(390))
}

func TestStadiumFeeCalculator_Calculate_For_BigFourWheeler(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	airportParkingCalculator := testFeeFactory.GetFeeCalculator(enum.Stadium)
	charges := airportParkingCalculator.Calculate(time.Now().Add(-3*time.Hour), enum.BigFourWheeler)
	assert.Equal(t, charges, float64(-1))
}
