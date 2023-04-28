package receipt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReceipt(t *testing.T) {
	mockDate := time.Date(2022, time.July, 15, 0, 0, 0, 0, time.Local)
	mockReceipt := NewReceipt(1, 1, mockDate, mockDate, 20)
	mockReceipt.Generate()
	assert.Equal(t, mockReceipt.GetParkingCharges(), float64(20))
}
