package main

import (
	"sahaj-parking-lot/clock"
	"sahaj-parking-lot/enum"
	"sahaj-parking-lot/parkinglot"
)

func main() {
	parkingLot := parkinglot.NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      3,
			enum.SmallFourWheeler: 1,
		}, enum.Airport, clock.NewClock())
	ticket1 := parkingLot.Park("Scooters")
	ticket1.Issue()
	ticket2 := parkingLot.Park("Scooters")
	ticket2.Issue()
	ticket3 := parkingLot.Park("Scooters")
	ticket3.Issue()
	receipt1 := parkingLot.UnPark(ticket2.GetNumber())
	receipt1.Generate()
	ticket4 := parkingLot.Park("Scooters")
	ticket4.Issue()
	ticket5 := parkingLot.Park("Cars")
	ticket5.Issue()
	parkingLot.UnPark(ticket1.GetNumber())
	receipt2 := parkingLot.UnPark(ticket5.GetNumber())
	receipt2.Generate()
	ticket6 := parkingLot.Park("Cars")
	ticket6.Issue()
	ticket7 := parkingLot.Park("Scooters")
	ticket7.Issue()
	receipt3 := parkingLot.UnPark(6)
	receipt3.Generate()
	receipt4 := parkingLot.UnPark(7)
	receipt4.Generate()

}
