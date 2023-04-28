package clock

import "time"

//go:generate mockgen -source=./clock.go -destination=./mocks/mock_clock.go -package=mocks
type Clock interface {
	Now() time.Time
}

type clock struct{}

func (c clock) Now() time.Time {
	location, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(location)
	return now
}

func NewClock() Clock {
	return &clock{}
}
