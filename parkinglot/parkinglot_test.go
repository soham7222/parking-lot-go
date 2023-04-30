package parkinglot

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sahaj-parking-lot/clock/mocks"
	"sahaj-parking-lot/enum"
	parkingErr "sahaj-parking-lot/error"
	"sahaj-parking-lot/feemodel"
	"testing"
	"time"
)

func Test_airport_parking_lot_calculate_parking_charges(t *testing.T) {
	mockController := gomock.NewController(t)
	mockClock := mocks.NewMockClock(mockController)

	tenHoursBeforeTime := time.Now().Add(-10 * time.Hour)
	sevenHoursBeforeTime := time.Now().Add(-7 * time.Hour)
	fiveHoursBeforeTime := time.Now().Add(-5 * time.Hour)
	oneHourBefore := time.Now().Add(-1 * time.Hour)
	tenMinutesBefore := time.Now().Add(-10 * time.Minute)
	fiveMinutesBefore := time.Now().Add(-5 * time.Minute)

	gomock.InOrder(
		//The first two Scooter have entered 10 hours before
		mockClock.EXPECT().Now().Return(tenHoursBeforeTime).Times(2),
		//The third scooter has entered 7 hours before
		mockClock.EXPECT().Now().Return(sevenHoursBeforeTime).Times(1),
		// One scooter gets un parked now
		mockClock.EXPECT().Now().Return(time.Now()).Times(1),
		// the fourth scooter have entered 5 hours before
		mockClock.EXPECT().Now().Return(fiveHoursBeforeTime).Times(1),
		// the first car has entered 1 hour before
		mockClock.EXPECT().Now().Return(oneHourBefore).Times(1),
		// one scooter and one car get un parked
		mockClock.EXPECT().Now().Return(time.Now()).Times(2),
		// second car has entered 10 minutes back
		mockClock.EXPECT().Now().Return(tenMinutesBefore).Times(1),
		// the fifth scooter have entered 10 minutes before
		mockClock.EXPECT().Now().Return(tenMinutesBefore).Times(1),
		// the sixth scooter have entered 10 minutes before
		mockClock.EXPECT().Now().Return(fiveMinutesBefore).Times(1),
		// two Scooter get un parked
		mockClock.EXPECT().Now().Return(time.Now()).Times(2),
		// the seventh scooter have entered 5 minutes before
		mockClock.EXPECT().Now().Return(fiveMinutesBefore).Times(1),
		// one truck have entered 5 minutes before
		mockClock.EXPECT().Now().Return(fiveMinutesBefore).Times(1),
	)

	feeFactory := feemodel.NewFeeFactory()
	parkingLot := NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      3,
			enum.SmallFourWheeler: 1,
		}, map[enum.SpotType]int{
			enum.TwoWheelers:      3,
			enum.SmallFourWheeler: 1,
		}, enum.Airport, mockClock, feeFactory)

	//first scooter parks to spot 0
	ticket1, err := parkingLot.Park("Scooter")
	assert.Equal(t, 0, ticket1.GetSpotNumber())
	assert.Equal(t, 1, ticket1.GetNumber())
	assert.Nil(t, err)

	//second scooter parks to spot 1
	ticket2, err := parkingLot.Park("Scooter")
	assert.Equal(t, 1, ticket2.GetSpotNumber())
	assert.Nil(t, err)

	//third scooter parks to spot 2
	ticket3, err := parkingLot.Park("Scooter")
	assert.Equal(t, 2, ticket3.GetSpotNumber())
	assert.Nil(t, err)

	//second scooter un parks from spot 1 & it gets charged 600 (10 hours * 60)
	receipt1, err := parkingLot.UnPark(ticket2.GetNumber())
	assert.Equal(t, float64(600), receipt1.GetParkingCharges())
	assert.Nil(t, err)

	//fourth scooter parks to spot 1 as spot 1 is empty and first available
	ticket4, err := parkingLot.Park("Scooter")
	ticket4.Issue()
	assert.Equal(t, 1, ticket4.GetSpotNumber())
	assert.Nil(t, err)

	// first car parks to spot 0 as it's different parking category from scooter/motorcycle
	ticket5, err := parkingLot.Park("Car")
	assert.Equal(t, 0, ticket5.GetSpotNumber())
	assert.Nil(t, err)

	// second scooter un parks from spot 2 & it gets charged 280 (7 hours * 40)
	receipt2, err := parkingLot.UnPark(ticket3.GetNumber())
	assert.Equal(t, float64(280), receipt2.GetParkingCharges())
	assert.Nil(t, err)

	// first car un parks from spot 0 & it gets charged 60 (1 hour * 40)
	receipt3, err := parkingLot.UnPark(ticket5.GetNumber())
	assert.Equal(t, float64(60), receipt3.GetParkingCharges())
	assert.Nil(t, err)

	// second car parks to spot 0
	ticket6, err := parkingLot.Park("Car")
	assert.Equal(t, 0, ticket6.GetSpotNumber())
	assert.Nil(t, err)

	//fifth scooter parks to spot 1 as spot 1 is empty and first available
	ticket7, err := parkingLot.Park("Scooter")
	assert.Equal(t, 2, ticket7.GetSpotNumber())
	assert.Nil(t, err)

	//sixth scooter comes for parking but parking is full so Ticket is nil
	ticket8, err := parkingLot.Park("Scooter")
	assert.Nil(t, ticket8)
	assert.Equal(t, err, parkingErr.ErrParkingFull)

	// fifth scooter un parks from spot 2 & it gets charged 0 (less than an hour)
	receipt4, err := parkingLot.UnPark(ticket7.GetNumber())
	assert.Equal(t, float64(0), receipt4.GetParkingCharges())
	assert.Nil(t, err)

	// the very first scooter un parks from spot 0 & it gets charged 600 (60 * 10 hours)
	receipt5, err := parkingLot.UnPark(ticket1.GetNumber())
	assert.Equal(t, float64(600), receipt5.GetParkingCharges())
	assert.Nil(t, err)

	//seventh scooter comes for parking with spot 0 and spot 2 opened. It will get parked in spot 0 as it's the first one
	ticket9, err := parkingLot.Park("Scooter")
	assert.Equal(t, 0, ticket9.GetSpotNumber())
	assert.Nil(t, err)

	// trying to un park from a ticket number which is not in the system should return nil
	receipt6, err := parkingLot.UnPark(99990)
	assert.Equal(t, nil, receipt6)
	assert.Equal(t, err, parkingErr.ErrInvalidTicketNumber)

	//Eighth scooter comes for parking but parking is full so Ticket is nil
	ticket10, err := parkingLot.Park("Truck")
	assert.Nil(t, ticket10)
	assert.Equal(t, err, parkingErr.ErrParkingNotSupported)
}
