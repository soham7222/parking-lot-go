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
	charges := mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.TwoWheelers)
	assert.Equal(t, charges, float64(20))
	charges = mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.BigFourWheeler)
	assert.Equal(t, charges, float64(100))
	charges = mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(40))
	charges = mallCalculator.Calculate(time.Now().Add(-2*time.Minute), enum.SmallFourWheeler)
	assert.Equal(t, charges, float64(20))
}

func TestMallFeeCalculator_Calculate_For_Unknown(t *testing.T) {
	testFeeFactory := NewFeeFactory()
	mallCalculator := testFeeFactory.GetFeeCalculator(enum.Mall)
	charges := mallCalculator.Calculate(time.Now().Add(-2*time.Hour), enum.UnknownType)
	assert.Equal(t, charges, float64(-1))
}
