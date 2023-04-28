package enum

type SpotType int

const (
	TwoWheelers SpotType = iota
	SmallFourWheeler
	BigFourWheeler
	UnknownType
)
