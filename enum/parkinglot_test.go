package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToParkingLotMode(t *testing.T) {
	assert.Equal(t, StringToParkingLotMode("Mall"), Mall)
	assert.Equal(t, StringToParkingLotMode("Stadium"), Stadium)
	assert.Equal(t, StringToParkingLotMode("Airport"), Airport)
}
