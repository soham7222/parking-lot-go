package spot

import (
	"sahaj-parking-lot/enum"
	"sahaj-parking-lot/utils"
	"time"
)

type Spot struct {
	vehicle          string
	occupied         bool
	number           int
	spotType         enum.SpotType
	vehicleEntryTime time.Time
}

var slotVehicleMapping = map[enum.SpotType]utils.Slice{
	enum.TwoWheelers:      {"Motorcycles", "Scooters"},
	enum.SmallFourWheeler: {"Cars", "SUVs"},
	enum.BigFourWheeler:   {"Buses", "Trucks"},
}

func (s *Spot) GetType() enum.SpotType {
	for slotType, vehicles := range slotVehicleMapping {
		if vehicles.Contains(s.vehicle) {
			s.spotType = slotType
			return slotType
		}
	}

	return enum.UnknownType
}

func (s *Spot) GetNumber() int {
	return s.number
}

func (s *Spot) IsOccupied() bool {
	return s.occupied
}

func (s *Spot) GetEntryTime() time.Time {
	return s.vehicleEntryTime
}

func (s *Spot) AssignSpotNumberAndMarkAsOccupied(spotNumber int) {
	s.occupied = true
	s.number = spotNumber
}

func NewSpot(vehicle string, vehicleEntryTime time.Time) Spot {
	return Spot{
		vehicle:          vehicle,
		vehicleEntryTime: vehicleEntryTime,
	}
}
