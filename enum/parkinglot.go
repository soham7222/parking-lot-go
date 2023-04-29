package enum

type ParkingLotType int

const (
	Mall ParkingLotType = iota
	Stadium
	Airport
)

var stringToParkingLotMode = map[string]ParkingLotType{
	"Mall":    1,
	"Stadium": 2,
	"Airport": 3,
}

func StringToParkingLotMode(spot string) ParkingLotType {
	return stringToParkingLotMode[spot]
}
