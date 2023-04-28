package spot

import (
	"sahaj-parking-lot/enum"
	"strings"
	"time"
)

type Spot struct {
	vehicle          string
	occupied         bool
	number           int
	spotType         enum.SpotType
	vehicleEntryTime time.Time
}

var slotVehicleMapping = map[enum.SpotType][]string{
	enum.TwoWheelers:      {"Motorcycle", "Scooter"},
	enum.SmallFourWheeler: {"Car", "SUV"},
	enum.BigFourWheeler:   {"Bus", "Truck"},
}

func (s *Spot) GetType() enum.SpotType {
	for slotType, vehicles := range slotVehicleMapping {
		if contains(vehicles, s.vehicle) {
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

func contains(source []string, value string) bool {
	for _, a := range source {
		if strings.EqualFold(a, value) {
			return true
		}
	}
	return false
}
