# Parking Lot problem

Given a parking lot with details about the vehicle types that can be parked, the number of spots, and the fee model for the parking lot; compute the fees to be paid for the parked vehicles when the vehicle is un parked.

### Run test locally with coverage
```
make test_local_with_coverage
```

### Run code locally
```
make install_deps
make run_local
```

### Run application locally
```
make build_local 
./sahaj-parking-lot
```

### Action available to run in parking lot console

* **Park**
  >It parks the car and issues a ticket.
* **Unpark**
  >it unparks the car and generates the receipt
* **exit**
  >it exits the console program
* **reset**
  >it resets the console and you can start from fresh

### Vehicle available for action as command

* Motorcycle
* Scooter
* Bus
* Truck
* Car
* SUV

#### Instruction
* use this following  command format to **Park** and **Unpark**
```
Park scooter
Unpark scooter, ticket number 001
```
>N.B: Please note that all the commands are case sensitive but vehicles are not

#### Example:

```sh
$ go run main.go
>
Commands available to use:
Type exit to exit the program
Type reset to reset to reset parking lot mode
Choose Your parking lot. Below options are available
1. Mall
2. Stadium
3. Airport
> Mall
Type the number of spots to be added for scooter/motorcycle 
> 3
Type the number of spots to be added for cars/suv
> 3
Type the number of spots to be added for bus/truck
> 0 
> Park scooter
------------------------------------------
Parking ticket:

ticket number: 001
spot number: 0
Entry Date: 30-Apr-2023 00:30:58
------------------------------------------
> Park car
------------------------------------------
Parking ticket:

ticket number: 002
spot number: 0
Entry Date: 30-Apr-2023 00:31:02
------------------------------------------
> Park SUV
------------------------------------------
Parking ticket:

ticket number: 003
spot number: 1
Entry Date: 30-Apr-2023 00:31:08
------------------------------------------
> Park motorcycle
------------------------------------------
Parking ticket:

ticket number: 004
spot number: 1
Entry Date: 30-Apr-2023 00:31:32
------------------------------------------
> Park car
------------------------------------------
Parking ticket:

ticket number: 005
spot number: 2
Entry Date: 30-Apr-2023 00:31:42
------------------------------------------
> Park car
The parking is full for car. Please come back later. 
> Park suv
The parking is full for suv. Please come back later. 
> Park scooter
------------------------------------------
Parking ticket:

ticket number: 006
spot number: 2
Entry Date: 30-Apr-2023 00:32:45
------------------------------------------
> Park scooter
parking is full
> Unpark scooter,ticket number 006       
------------------------------------------
Parking receipt:

receipt number: R-001
Ticket number: 006
Entry Date-time: 30-Apr-2023 00:32:45
Exit Date-time: 30-Apr-2023 00:33:55
Fees: 10
------------------------------------------
> Park scooter
------------------------------------------
Parking ticket:

ticket number: 007
spot number: 2
Entry Date: 30-Apr-2023 00:34:07
------------------------------------------
> Unpark scooter,ticket number 005
------------------------------------------
Parking receipt:

receipt number: R-002
Ticket number: 005
Entry Date-time: 30-Apr-2023 00:31:42
Exit Date-time: 30-Apr-2023 00:34:35
Fees: 20
------------------------------------------
> Park SUV
------------------------------------------
Parking ticket:

ticket number: 008
spot number: 2
Entry Date: 30-Apr-2023 00:34:45
------------------------------------------ 
> Park Truck
parking not supported
> exit
```

#### Test Cases Added for following scenario:

Mall parking lot
Spots:
● Motorcycles/scooters: 100 spots
● Cars/SUVs: 80 spots
● Buses/Trucks: 10 spots
Fee Model: Please refer to the Mall fee model and its examples, mentioned in the ‘Fee
Models’ section
Scenarios: The park and unpark steps shown in the previous example have been skipped to
reduce the text in the problem statement.
● Motorcycle parked for 3 hours and 30 mins. Fees: 40
● Car parked for 6 hours and 1 min. Fees: 140
● Truck parked for 1 hour and 59 mins. Fees: 100

Stadium Parking Lot
Spots:
● Motorcycles/scooters: 1000 spots
● Cars/SUVs: 1500 spots
Fee Model: Please refer to the Stadium fee model mentioned in the ‘Fee Models’ section
Scenarios: The park and unpark steps shown in the previous example have been skipped to
reduce the text in the problem statement.
● Motorcycle parked for 3 hours and 40 mins. Fees: 30
● Motorcycle parked for 14 hours and 59 mins. Fees: 390.
○ 30 for the first 4 hours. 60 for the next 8 hours. And then 300 for the
remaining duration.
● Electric SUV parked for 11 hours and 30 mins. Fees: 180.
○ 60 for the first 4 hours and then 120 for the remaining duration.
● SUV parked for 13 hours and 5 mins. Fees: 580.
○ 60 for the first 4 hours and then 120 for the next 8 hours. 400 for the
remaining duration.

Airport Parking Lot
Spots:
● Motorcycles/scooters: 200 spots
● Cars/SUVs: 500 spots
● Buses/Trucks: 100 spots
Fee Model: Please refer to the Airport fee model mentioned in the ‘Fee Models’ section
Scenarios: The park and unpark steps shown in the previous example have been skipped to
reduce the text in the problem statement.
5
● Motorcycle parked for 55 mins. Fees: 0
● Motorcycle parked for 14 hours and 59 mins. Fees: 60
● Motorcycle parked for 1 day and 12 hours. Fees: 160
● Car parked for 50 mins. Fees: 60
● SUV parked for 23 hours and 59 mins. Fees: 80
● Car parked for 3 days and 1 hour. Fees: 400
