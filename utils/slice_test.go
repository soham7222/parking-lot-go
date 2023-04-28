package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_Contains(t *testing.T) {
	stringTypeSource := Slice{"Jam", "Jelly", "Yogurt"}
	assert.Equal(t, true, stringTypeSource.Contains("Jelly"))
	assert.Equal(t, true, stringTypeSource.Contains("Jam"))
	assert.Equal(t, true, stringTypeSource.Contains("Yogurt"))
	assert.Equal(t, false, stringTypeSource.Contains("Apple"))

	intTypeSource := Slice{1, 2, 3}
	assert.Equal(t, true, intTypeSource.Contains(1))
	assert.Equal(t, true, intTypeSource.Contains(2))
	assert.Equal(t, true, intTypeSource.Contains(3))
	assert.Equal(t, false, intTypeSource.Contains(99))
}
