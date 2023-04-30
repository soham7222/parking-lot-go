package enum

type ParkingLotType int

const (
	Mall ParkingLotType = iota
	Stadium
	Airport
)

var stringToParkingLotMode = map[string]ParkingLotType{
	"Mall":    0,
	"Stadium": 1,
	"Airport": 2,
}

func StringToParkingLotMode(spot string) ParkingLotType {
	return stringToParkingLotMode[spot]
}
