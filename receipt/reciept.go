package receipt

import (
	"fmt"
	"time"
)

type Receipt interface {
	Generate()
	GetParkingCharges() float64
}

type receipt struct {
	id               string
	ticketNumber     int
	vehicleEntryTime time.Time
	vehicleExitTime  time.Time
	parkingCharges   float64
}

func (r *receipt) Generate() {

	fmt.Println("------------------------------------------")
	fmt.Println("Parking receipt:")
	fmt.Println("")
	fmt.Println("receipt number:", r.id)
	fmt.Println("Ticket number:", fmt.Sprintf("%03d", r.ticketNumber))
	fmt.Println("Entry Date-time:", r.vehicleEntryTime.Format("02-Jan-2006 15:04:05"))
	fmt.Println("Exit Date-time:", r.vehicleExitTime.Format("02-Jan-2006 15:04:05"))
	fmt.Println("Fees:", r.parkingCharges)
	fmt.Println("------------------------------------------")
}

func (r *receipt) GetParkingCharges() float64 {
	return r.parkingCharges
}

func NewReceipt(id,
	ticketNumber int,
	vehicleEntryTime time.Time,
	vehicleExitTime time.Time,
	parkingCharges float64,
) Receipt {
	return &receipt{
		"R-" + fmt.Sprintf("%03d", id),
		ticketNumber,
		vehicleEntryTime,
		vehicleExitTime,
		parkingCharges,
	}
}
