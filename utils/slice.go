package utils

type Slice []any

func (e Slice) Contains(value any) bool {
	for _, a := range e {
		if a == value {
			return true
		}
	}
	return false
}
