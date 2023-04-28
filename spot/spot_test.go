package spot

import (
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/enum"
	"testing"
	"time"
)

func TestSpot(t *testing.T) {
	mockDate := time.Date(2022, time.July, 15, 0, 0, 0, 0, time.Local)
	mockSpot := NewSpot("Car", mockDate)
	assert.Equal(t, mockSpot.IsOccupied(), false)
	mockSpot.AssignSpotNumberAndMarkAsOccupied(12)
	assert.Equal(t, mockSpot.GetNumber(), 12)
	assert.Equal(t, mockSpot.GetType(), enum.SmallFourWheeler)
	assert.Equal(t, mockSpot.GetEntryTime(), mockDate)
	assert.Equal(t, mockSpot.IsOccupied(), true)

	mockSpot2 := NewSpot("test", mockDate)
	assert.Equal(t, mockSpot2.GetType(), enum.UnknownType)
}
