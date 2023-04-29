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

### Action available to run in parking lot console

* **Park**
  >It parks the car and issues a ticket.
* **Unpark**
  >it unparks the car and generates the receipt
* **exit**
  >it exits the console program

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
The parking is full for scooter. Please come back later. 
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
> exit
```
