package ticket

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/clock/mocks"
	"testing"
	"time"
)

func TestNewTicket(t *testing.T) {
	mockController := gomock.NewController(t)
	mockClock := mocks.NewMockClock(mockController)
	mockDate := time.Date(2022, time.July, 15, 0, 0, 0, 0, time.Local)
	mockClock.EXPECT().Now().Return(mockDate)
	tckt := NewTicket("12", 10, mockClock.Now())
	tckt.Issue()
	assert.Equal(t, tckt.GetNumber(), "12")
	assert.Equal(t, tckt.GetSpotNumber(), 10)
	assert.Equal(t, tckt.GetEntryTime(), mockDate)
}
