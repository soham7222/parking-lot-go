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

func Test_airport_parking_lot_spot_management(t *testing.T) {
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
	assert.Equal(t, float64(60), receipt1.GetParkingCharges())
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
	assert.Equal(t, float64(40), receipt2.GetParkingCharges())
	assert.Nil(t, err)

	// first car un parks from spot 0 & it gets charged 60 (1 hour * 60)
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
	assert.Equal(t, float64(60), receipt5.GetParkingCharges())
	assert.Nil(t, err)

	//seventh scooter comes for parking with spot 0 and spot 2 opened. It will get parked in spot 0 as it's the first one
	ticket9, err := parkingLot.Park("Scooter")
	assert.Equal(t, 0, ticket9.GetSpotNumber())
	assert.Nil(t, err)

	// trying to un park from a ticket number which is not in the system should return nil
	receipt6, err := parkingLot.UnPark(99990)
	assert.Nil(t, receipt6)
	assert.Equal(t, err, parkingErr.ErrInvalidTicketNumber)

	//Eighth scooter comes for parking but parking is full so Ticket is nil
	ticket10, err := parkingLot.Park("Truck")
	assert.Nil(t, ticket10)
	assert.Equal(t, err, parkingErr.ErrParkingNotSupported)

	// trying to generate a receipt of already left vehicle
	receipt7, err := parkingLot.UnPark(ticket1.GetNumber())
	assert.Equal(t, receipt5, receipt7)
}

//Example 2: Mall parking lot
//Spots:
//● Motorcycles/scooters: 100 spots
//● Cars/SUVs: 80 spots
//● Buses/Trucks: 10 spots
//Fee Model: Please refer to the Mall fee model and its examples, mentioned in the ‘Fee
//Models’ section
//Scenarios: The park and unpark steps shown in the previous example have been skipped to
//reduce the text in the problem statement.
//● Motorcycle parked for 3 hours and 30 mins. Fees: 40
//● Car parked for 6 hours and 1 min. Fees: 140
//● Truck parked for 1 hour and 59 mins. Fees: 100
func Test_Mall_parking_lot_calculate_parking_charges(t *testing.T) {
	mockController := gomock.NewController(t)
	mockClock := mocks.NewMockClock(mockController)

	gomock.InOrder(
		//● Motorcycle parked for 3 hours and 30 mins.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-3*time.Hour).Add(-30*time.Minute)).Times(1),
		//● Car parked for 6 hours and 1 min.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-6*time.Hour).Add(-1*time.Minute)).Times(1),
		//● Truck parked for 1 hour and 59 mins.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-1*time.Hour).Add(-59*time.Minute)).Times(1),
		// All three vehicles are getting un parked
		mockClock.EXPECT().Now().Return(time.Now()).Times(3),
	)

	feeFactory := feemodel.NewFeeFactory()
	parkingLot := NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      100,
			enum.SmallFourWheeler: 80,
			enum.BigFourWheeler:   10,
		}, map[enum.SpotType]int{
			enum.TwoWheelers:      100,
			enum.SmallFourWheeler: 80,
			enum.BigFourWheeler:   10,
		}, enum.Mall, mockClock, feeFactory)

	//● Motorcycle parked for 3 hours and 30 mins. Fees: 40
	//● Car parked for 6 hours and 1 min. Fees: 140
	//● Truck parked for 1 hour and 59 mins. Fees: 100
	ticket1, _ := parkingLot.Park("Scooter")
	ticket2, _ := parkingLot.Park("Car")
	ticket3, _ := parkingLot.Park("Truck")

	//Motorcycle parked for 3 hours and 40 mins. Fees: 40
	receipt1, err := parkingLot.UnPark(ticket1.GetNumber())
	assert.Equal(t, float64(40), receipt1.GetParkingCharges())
	assert.Nil(t, err)

	//Car parked for 6 hours and 1 min. Fees: 140
	receipt2, err := parkingLot.UnPark(ticket2.GetNumber())
	assert.Equal(t, float64(140), receipt2.GetParkingCharges())
	assert.Nil(t, err)

	//● Truck parked for 1 hour and 59 mins. Fees: 100
	receipt3, err := parkingLot.UnPark(ticket3.GetNumber())
	assert.Equal(t, float64(100), receipt3.GetParkingCharges())
	assert.Nil(t, err)
}

//Example 3: Stadium Parking Lot
//Spots:
//● Motorcycles/scooters: 1000 spots
//● Cars/SUVs: 1500 spots
//Fee Model: Please refer to the Stadium fee model mentioned in the ‘Fee Models’ section
//Scenarios: The park and unpark steps shown in the previous example have been skipped to
//reduce the text in the problem statement.
//● Motorcycle parked for 3 hours and 40 mins. Fees: 30
//● Motorcycle parked for 14 hours and 59 mins. Fees: 390.
//○ 30 for the first 4 hours. 60 for the next 8 hours. And then 300 for the
//remaining duration.
//● Electric SUV parked for 11 hours and 30 mins. Fees: 180.
//○ 60 for the first 4 hours and then 120 for the remaining duration.
//● SUV parked for 13 hours and 5 mins. Fees: 580.
//○ 60 for the first 4 hours and then 120 for the next 8 hours. 400 for the
//remaining duration.
func Test_Stadium_parking_lot_calculate_parking_charges(t *testing.T) {
	mockController := gomock.NewController(t)
	mockClock := mocks.NewMockClock(mockController)

	gomock.InOrder(
		//Motorcycle parked for 3 hours and 40 mins
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-3*time.Hour).Add(-40*time.Minute)).Times(1),
		//● Motorcycle parked for 14 hours and 59 mins.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-14*time.Hour).Add(-59*time.Minute)).Times(1),
		//● Electric SUV parked for 11 hours and 30 mins
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-11*time.Hour).Add(-30*time.Minute)).Times(1),
		//○ 60 for the first 4 hours and then 120 for the next 8 hours.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-13*time.Hour).Add(-59*time.Minute)).Times(1),
		// All four vehicles are getting un parked
		mockClock.EXPECT().Now().Return(time.Now()).Times(4),
	)

	feeFactory := feemodel.NewFeeFactory()
	parkingLot := NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      1000,
			enum.SmallFourWheeler: 1500,
		}, map[enum.SpotType]int{
			enum.TwoWheelers:      1000,
			enum.SmallFourWheeler: 1500,
		}, enum.Stadium, mockClock, feeFactory)

	//first motorcycle parks to spot 0
	ticket1, err := parkingLot.Park("Scooter")
	assert.Equal(t, 0, ticket1.GetSpotNumber())
	assert.Equal(t, 1, ticket1.GetNumber())
	assert.Nil(t, err)

	//second motorcycle parks to spot 1
	ticket2, err := parkingLot.Park("Scooter")
	assert.Equal(t, 1, ticket2.GetSpotNumber())
	assert.Nil(t, err)

	// first SUV parks to spot 0 as it's different parking category from scooter/motorcycle
	ticket3, err := parkingLot.Park("Car")
	assert.Equal(t, 0, ticket3.GetSpotNumber())
	assert.Nil(t, err)

	// second SUV parks to spot 1
	ticket4, err := parkingLot.Park("Car")
	assert.Equal(t, 1, ticket4.GetSpotNumber())
	assert.Nil(t, err)

	//Motorcycle parked for 3 hours and 40 mins. Fees: 30
	receipt1, err := parkingLot.UnPark(ticket1.GetNumber())
	assert.Equal(t, float64(30), receipt1.GetParkingCharges())
	assert.Nil(t, err)

	//● Motorcycle parked for 14 hours and 59 mins. Fees: 390.
	//○ 30 for the first 4 hours. 60 for the next 8 hours. And then 300 for the
	//remaining duration.
	receipt2, err := parkingLot.UnPark(ticket2.GetNumber())
	assert.Equal(t, float64(390), receipt2.GetParkingCharges())
	assert.Nil(t, err)

	//● Electric SUV parked for 11 hours and 30 mins. Fees: 180.
	//○ 60 for the first 4 hours and then 120 for the remaining duration.
	receipt3, err := parkingLot.UnPark(ticket3.GetNumber())
	assert.Equal(t, float64(180), receipt3.GetParkingCharges())
	assert.Nil(t, err)

	//● SUV parked for 13 hours and 5 mins. Fees: 580.
	//○ 60 for the first 4 hours and then 120 for the next 8 hours. 400 for the
	//remaining duration.
	receipt4, err := parkingLot.UnPark(ticket4.GetNumber())
	assert.Equal(t, float64(580), receipt4.GetParkingCharges())
	assert.Nil(t, err)
}

//Example 4: Airport Parking Lot
//Spots:
//● Motorcycles/scooters: 200 spots
//● Cars/SUVs: 500 spots
//● Buses/Trucks: 100 spots
//Fee Model: Please refer to the Airport fee model mentioned in the ‘Fee Models’ section
//Scenarios: The park and unpark steps shown in the previous example have been skipped to
//reduce the text in the problem statement.
//5
//● Motorcycle parked for 55 mins. Fees: 0
//● Motorcycle parked for 14 hours and 59 mins. Fees: 60
//● Motorcycle parked for 1 day and 12 hours. Fees: 160
//● Car parked for 50 mins. Fees: 60
//● SUV parked for 23 hours and 59 mins. Fees: 80
//● Car parked for 3 days and 1 hour. Fees: 400
func Test_airport_parking_lot_calculate_parking_charges(t *testing.T) {
	mockController := gomock.NewController(t)
	mockClock := mocks.NewMockClock(mockController)

	gomock.InOrder(
		//Motorcycle parked for 55 mins.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-55*time.Minute)).Times(1),
		//● Motorcycle parked for 14 hours and 59 mins.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-14*time.Hour).Add(-59*time.Minute)).Times(1),
		//● Motorcycle parked for 1day 12 hours.
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-36*time.Hour)).Times(1),
		//Car parked for 50 mins. Fees: 60
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-50*time.Minute)).Times(1),
		//● ● SUV parked for 23 hours and 59 mins. Fees: 80
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-23*time.Hour).Add(-59*time.Minute)).Times(1),
		//● Car parked for 3 days and 1 hour. Fees: 400
		mockClock.EXPECT().Now().Return(
			time.Now().Add(-73*time.Hour)).Times(1),
		// All four vehicles are getting un parked
		mockClock.EXPECT().Now().Return(time.Now()).Times(6),
	)

	feeFactory := feemodel.NewFeeFactory()
	parkingLot := NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      200,
			enum.SmallFourWheeler: 500,
			enum.BigFourWheeler:   100,
		}, map[enum.SpotType]int{
			enum.TwoWheelers:      200,
			enum.SmallFourWheeler: 500,
			enum.BigFourWheeler:   100,
		}, enum.Airport, mockClock, feeFactory)

	ticket1, err := parkingLot.Park("Motorcycle")
	assert.Equal(t, 0, ticket1.GetSpotNumber())
	assert.Equal(t, 1, ticket1.GetNumber())
	assert.Nil(t, err)

	ticket2, err := parkingLot.Park("Motorcycle")
	assert.Equal(t, 1, ticket2.GetSpotNumber())
	assert.Nil(t, err)

	ticket3, err := parkingLot.Park("Motorcycle")
	assert.Equal(t, 2, ticket3.GetSpotNumber())
	assert.Nil(t, err)

	ticket4, err := parkingLot.Park("Car")
	assert.Equal(t, 0, ticket4.GetSpotNumber())
	assert.Nil(t, err)

	ticket5, err := parkingLot.Park("Suv")
	assert.Equal(t, 1, ticket5.GetSpotNumber())
	assert.Nil(t, err)

	ticket6, err := parkingLot.Park("Suv")
	assert.Equal(t, 2, ticket6.GetSpotNumber())
	assert.Nil(t, err)

	receipt1, err := parkingLot.UnPark(ticket1.GetNumber())
	assert.Equal(t, float64(0), receipt1.GetParkingCharges())
	assert.Nil(t, err)

	receipt2, err := parkingLot.UnPark(ticket2.GetNumber())
	assert.Equal(t, float64(60), receipt2.GetParkingCharges())
	assert.Nil(t, err)

	receipt3, err := parkingLot.UnPark(ticket3.GetNumber())
	assert.Equal(t, float64(160), receipt3.GetParkingCharges())
	assert.Nil(t, err)

	receipt4, err := parkingLot.UnPark(ticket4.GetNumber())
	assert.Equal(t, float64(60), receipt4.GetParkingCharges())
	assert.Nil(t, err)

	receipt5, err := parkingLot.UnPark(ticket5.GetNumber())
	assert.Equal(t, float64(80), receipt5.GetParkingCharges())
	assert.Nil(t, err)

	receipt6, err := parkingLot.UnPark(ticket6.GetNumber())
	assert.Equal(t, float64(400), receipt6.GetParkingCharges())
	assert.Nil(t, err)
}
