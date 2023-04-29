package error

import "errors"

var ErrParkingFull = errors.New("parking is full")
var ErrParkingNotSupported = errors.New("parking not supported")
var ErrInvalidTicketNumber = errors.New("in valid ticket number")
