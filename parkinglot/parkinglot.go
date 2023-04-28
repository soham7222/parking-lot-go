package parkinglot

import (
	"fmt"
	"sahaj-parking-lot/clock"
	"sahaj-parking-lot/enum"
	"sahaj-parking-lot/feemodel"
	"sahaj-parking-lot/receipt"
	"sahaj-parking-lot/spot"
	"sahaj-parking-lot/ticket"
)

type ParkingLot interface {
	Park(vehicle string) ticket.Ticket
	UnPark(ticketNumber int) receipt.Receipt
}

type parkingLot struct {
	totalCapacity      map[enum.SpotType]int
	spots              map[enum.SpotType][]spot.Spot
	remainingCapacity  map[enum.SpotType]int
	parkingType        enum.ParkingLotType
	ticketSpotMapper   map[int]*spot.Spot
	ticketIdGenerator  func() int
	receiptIdGenerator func() int
	clock              clock.Clock
}

func NewParkingLot(totalCapacity map[enum.SpotType]int,
	parkingType enum.ParkingLotType, clock clock.Clock) ParkingLot {
	return &parkingLot{
		totalCapacity:      totalCapacity,
		remainingCapacity:  totalCapacity,
		spots:              getSpots(totalCapacity),
		parkingType:        parkingType,
		ticketSpotMapper:   make(map[int]*spot.Spot, 0),
		ticketIdGenerator:  counter(),
		receiptIdGenerator: counter(),
		clock:              clock,
	}
}

func (p *parkingLot) Park(vehicle string) ticket.Ticket {
	entryTime := p.clock.Now()
	assignedSpot := spot.NewSpot(vehicle, entryTime)
	assignedSpotType := assignedSpot.GetType()

	assignedSlotNumber := p.getFirstAvailableSlotNumber(assignedSpotType)
	if assignedSlotNumber < 0 {
		fmt.Println("full")
		return nil
	} else {
		assignedSpot.AssignSpotNumberAndMarkAsOccupied(assignedSlotNumber)
		spotTypeSlots := p.spots[assignedSpotType]
		spotTypeSlots[assignedSlotNumber] = assignedSpot

		p.remainingCapacity[assignedSpotType]--

		ticketNumber := p.ticketIdGenerator()
		p.ticketSpotMapper[ticketNumber] = &assignedSpot
		ticketToBeIssued := ticket.NewTicket(ticketNumber, assignedSpot.GetNumber(), entryTime)
		return ticketToBeIssued
	}
}

func (p *parkingLot) UnPark(ticketNumber int) receipt.Receipt {
	slotToBeFreed := p.ticketSpotMapper[ticketNumber]

	if slotToBeFreed != nil {
		p.remainingCapacity[slotToBeFreed.GetType()]++

		p.spots[slotToBeFreed.GetType()][slotToBeFreed.GetNumber()] = spot.Spot{}

		feeFactory := feemodel.NewFeeFactory()
		feeCalculator := feeFactory.GetFeeCalculator(p.parkingType)
		parkingCharges := feeCalculator.Calculate(slotToBeFreed.GetEntryTime(), slotToBeFreed.GetType())
		receiptForParking := receipt.NewReceipt(p.receiptIdGenerator(),
			ticketNumber, slotToBeFreed.GetEntryTime(), p.clock.Now(), parkingCharges)
		return receiptForParking
	} else {
		return nil
	}
}

func (p *parkingLot) getFirstAvailableSlotNumber(spotType enum.SpotType) int {
	if p.remainingCapacity[spotType] > 0 {
		for i := 0; i <= p.totalCapacity[spotType]+1; i++ {
			if !p.spots[spotType][i].IsOccupied() {
				return i
			}
		}
	}
	return -1
}

func getSpots(totalCapacity map[enum.SpotType]int) map[enum.SpotType][]spot.Spot {
	result := make(map[enum.SpotType][]spot.Spot)
	for key, value := range totalCapacity {
		result[key] = make([]spot.Spot, value)
	}
	return result
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
