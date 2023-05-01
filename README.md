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
$ make run_local
go run main.go
Commands available to use:
Type exit to exit the program
Type exit to reset to reset parking lot mode
Choose Your parking lot. Below options are available
1. Mall
2. Stadium
3. Airport
> Mall
Type the number of spots to be added for scooter/motorcycle 
> 2
Type the number of spots to be added for cars/suv
> 2
Type the number of spots to be added for bus/truck
> 0
Now you can start parking and unparking scooter, motorcycle, cars, suv, bus and truck
Use the below command structure to park and unpark vehicles
Park Scooter
Unpark Scooter, ticket number 001

> Park Scooter
------------------------------------------
Parking ticket:

ticket number: 001
spot number: 0
Entry Date: 01-May-2023 13:19:51
------------------------------------------
> Park Car
------------------------------------------
Parking ticket:

ticket number: 002
spot number: 0
Entry Date: 01-May-2023 13:19:54
------------------------------------------
> Park motorcycle
------------------------------------------
Parking ticket:

ticket number: 003
spot number: 1
Entry Date: 01-May-2023 13:20:02
------------------------------------------
> Park SUV
------------------------------------------
Parking ticket:

ticket number: 004
spot number: 1
Entry Date: 01-May-2023 13:20:10
------------------------------------------
> Park Scooter
parking is full
> Park SUV
parking is full
> Unpark Scooter, ticket number 001
------------------------------------------
Parking receipt:

receipt number: R-001
Ticket number: 001
Entry Date-time: 01-May-2023 13:19:51
Exit Date-time: 01-May-2023 13:21:19
Fees: 10
------------------------------------------
> Park Scooter
------------------------------------------
Parking ticket:

ticket number: 005
spot number: 0
Entry Date: 01-May-2023 13:21:29
------------------------------------------
> Park Bus
parking not supported
> Unpark Scooter, ticket number 001
the vehicle you are trying to unpark has already left the premised. Here is the receipt for that
------------------------------------------
Parking receipt:

receipt number: R-001
Ticket number: 001
Entry Date-time: 01-May-2023 13:19:51
Exit Date-time: 01-May-2023 13:21:19
Fees: 10
------------------------------------------
> exit
```