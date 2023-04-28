package ticket

import (
	"fmt"
	"time"
)

type Ticket interface {
	Issue()
	GetEntryTime() time.Time
	GetNumber() int
	GetSpotNumber() int
}

type ticket struct {
	number           int
	spotNumber       int
	vehicleEntryTime time.Time
}

func NewTicket(number int,
	spotNumber int,
	vehicleEntryTime time.Time) Ticket {
	return &ticket{
		number:           number,
		spotNumber:       spotNumber,
		vehicleEntryTime: vehicleEntryTime,
	}
}

func (t *ticket) Issue() {
	fmt.Println("------------------------------------------")
	fmt.Println("Parking ticket:")
	fmt.Println("")
	fmt.Println("ticket number:", t.number)
	fmt.Println("spot number:", t.spotNumber)
	fmt.Println("Entry Date:", t.vehicleEntryTime.Format("02-Jan-2006 15:04:05"))
	fmt.Println("------------------------------------------")
}

func (t *ticket) GetEntryTime() time.Time {
	return t.vehicleEntryTime
}

func (t *ticket) GetNumber() int {
	return t.number
}

func (t *ticket) GetSpotNumber() int {
	return t.spotNumber
}
