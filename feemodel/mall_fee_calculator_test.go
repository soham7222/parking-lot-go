package feemodel

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestMallFeeCalculator_Calculate(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	mallCalculator := testFeeFactory.GetFeeCalculator(enum.Mall)
	charges := mallCalculator.Calculate(
		time.Now().Add(-3*time.Hour).Add(-30*time.Minute), enum.TwoWheelers)
	assert.Equal(t, charges, float64(40))
	charges = mallCalculator.Calculate(
		time.Now().Add(-1*time.Hour).Add(-59*time.Minute), enum.BigFourWheeler)
	assert.Equal(t, charges, float64(100))
	charges = mallCalculator.Calculate(
		time.Now().Add(-6*time.Hour).Add(-1*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(140))
	charges = mallCalculator.Calculate(time.Now().Add(-2*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(20))
}

func TestMallFeeCalculator_Calculate_For_Unknown(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	mallCalculator := testFeeFactory.GetFeeCalculator(enum.Mall)
	charges := mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.UnknownType)
	assert.Equal(t, charges, float64(-1))
}
