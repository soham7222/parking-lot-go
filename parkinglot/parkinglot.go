package parkinglot

import (
	"fmt"
	"sahaj-parking-lot/clock"
	"sahaj-parking-lot/enum"
	parkingErr "sahaj-parking-lot/error"
	"sahaj-parking-lot/feemodel"
	"sahaj-parking-lot/receipt"
	"sahaj-parking-lot/spot"
	"sahaj-parking-lot/ticket"
)

type ParkingLot interface {
	Park(vehicle string) (ticket.Ticket, error)
	UnPark(ticketNumber int) (receipt.Receipt, error)
}

type parkingLot struct {
	totalCapacity      map[enum.SpotType]int
	remainingCapacity  map[enum.SpotType]int
	spots              map[enum.SpotType][]spot.Spot
	parkingType        enum.ParkingLotType
	ticketSpotMapper   map[int]*spot.Spot
	ticketIdGenerator  func() int
	receiptIdGenerator func() int
	clock              clock.Clock
	feeFactory         feemodel.FeeFactory
}

func NewParkingLot(totalCapacity map[enum.SpotType]int,
	remainingCapacity map[enum.SpotType]int,
	parkingType enum.ParkingLotType, clock clock.Clock,
	feeFactory feemodel.FeeFactory) ParkingLot {
	return &parkingLot{
		totalCapacity:      totalCapacity,
		remainingCapacity:  totalCapacity,
		spots:              initializeSpots(totalCapacity),
		parkingType:        parkingType,
		ticketSpotMapper:   make(map[int]*spot.Spot, 0),
		ticketIdGenerator:  counter(),
		receiptIdGenerator: counter(),
		clock:              clock,
		feeFactory:         feeFactory,
	}
}

func (p *parkingLot) Park(vehicle string) (ticket.Ticket, error) {
	entryTime := p.clock.Now()
	assignedSpot := spot.NewSpot(vehicle, entryTime)
	assignedSpotType := assignedSpot.GetType()

	if len(p.spots[assignedSpotType]) == 0 {
		return nil, parkingErr.ErrParkingNotSupported
	}

	assignedSlotNumber := p.getFirstAvailableSlotNumber(assignedSpotType)
	if assignedSlotNumber < 0 {
		return nil, parkingErr.ErrParkingFull
	} else {
		assignedSpot.AssignSpotNumberAndMarkAsOccupied(assignedSlotNumber)
		spotTypeSlots := p.spots[assignedSpotType]
		spotTypeSlots[assignedSlotNumber] = assignedSpot

		p.remainingCapacity[assignedSpotType]--

		ticketNumber := p.ticketIdGenerator()
		p.ticketSpotMapper[ticketNumber] = &assignedSpot
		ticketToBeIssued := ticket.NewTicket(ticketNumber, assignedSpot.GetNumber(), entryTime)
		return ticketToBeIssued, nil
	}
}

func (p *parkingLot) UnPark(ticketNumber int) (receipt.Receipt, error) {
	slotToBeFreed := p.ticketSpotMapper[ticketNumber]

	if slotToBeFreed != nil {
		p.remainingCapacity[slotToBeFreed.GetType()]++

		p.spots[slotToBeFreed.GetType()][slotToBeFreed.GetNumber()] = spot.Spot{}

		parkingCharges := p.feeFactory.
			GetFeeCalculator(p.parkingType).
			Calculate(slotToBeFreed.GetVehicleEntryTime(), slotToBeFreed.GetType())

		receiptForParking := receipt.NewReceipt(p.receiptIdGenerator(),
			ticketNumber, slotToBeFreed.GetVehicleEntryTime(), p.clock.Now(), parkingCharges)
		return receiptForParking, nil
	} else {
		return nil, parkingErr.ErrInvalidTicketNumber
	}
}

func (p *parkingLot) getFirstAvailableSlotNumber(spotType enum.SpotType) int {
	fmt.Println(p.totalCapacity[spotType])
	if p.remainingCapacity[spotType] > 0 {
		for i := 0; i <= p.totalCapacity[spotType]+1; i++ {
			if !p.spots[spotType][i].IsOccupied() {
				return i
			}
		}
	}
	return -1
}

func initializeSpots(totalCapacity map[enum.SpotType]int) map[enum.SpotType][]spot.Spot {
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
