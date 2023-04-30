package main

import (
	"bufio"
	"fmt"
	"os"
	"sahaj-parking-lot/clock"
	"sahaj-parking-lot/enum"
	"sahaj-parking-lot/feemodel"
	"sahaj-parking-lot/parkinglot"
	"sahaj-parking-lot/spot"
	"strconv"
	"strings"
)

var (
	feeFactory              = feemodel.NewFeeFactory()
	parkingLot              parkinglot.ParkingLot
	parkingLotMode          string
	twoWheelerParking       string
	smallFourWheelerParking string
	bigFourWheelerParking   string
)

func main() {
	Init()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "reset" {
			Init()
		}

		if input == "exit" {
			break
		}

		commands := strings.Split(input, " ")
		if len(commands) > 1 {
			if contains(commandAllowed(), strings.TrimSpace(commands[0])) {
				if strings.Contains(input, "Unpark") {
					initiateUnParkAndGenerateReceipt(commands)
				} else {
					if contains(vehiclesAllowed(), strings.TrimSpace(commands[1])) {
						initiateParkAndIssueTicket(commands)
					} else {
						fmt.Printf("currently parking is not available for %s type of vehicle", commands[1])
					}
				}
			} else {
				fmt.Println("Not a valid action. Kindly read the documentation")
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func Init() {
	fmt.Println("Commands available to use:")
	fmt.Println("Type exit to exit the program")
	fmt.Println("Type exit to reset to reset parking lot mode")
	fmt.Println("Choose Your parking lot. Below options are available")
	fmt.Println("1. Mall\n2. Stadium\n3. Airport")
	fmt.Print("> ")
	fmt.Scanln(&parkingLotMode)
	fmt.Println("Type the number of spots to be added for scooter/motorcycle ")
	fmt.Print("> ")
	fmt.Scanln(&twoWheelerParking)
	fmt.Println("Type the number of spots to be added for cars/suv")
	fmt.Print("> ")
	fmt.Scanln(&smallFourWheelerParking)
	fmt.Println("Type the number of spots to be added for bus/truck")
	fmt.Print("> ")
	fmt.Scanln(&bigFourWheelerParking)
	twoWheelerParkingCount, _ := strconv.Atoi(twoWheelerParking)
	smallFourWheelerParkingCount, _ := strconv.Atoi(smallFourWheelerParking)
	bigFourWheelerParkingCount, _ := strconv.Atoi(bigFourWheelerParking)
	parkingLot = parkinglot.NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      twoWheelerParkingCount,
			enum.SmallFourWheeler: smallFourWheelerParkingCount,
			enum.BigFourWheeler:   bigFourWheelerParkingCount,
		}, enum.StringToParkingLotMode(parkingLotMode),
		clock.NewClock(), feeFactory)
	fmt.Println("Now you can start parking and unparking scooter, motorcycle, cars, suv, bus and truck")
	fmt.Println("Use the below command structure to park and unpark vehicles")
	fmt.Println("Park Scooter")
	fmt.Println("Unpark Scooter, ticket number 001")
	fmt.Println("")
}

func initiateParkAndIssueTicket(commands []string) {
	ticket, err := parkingLot.Park(commands[1])
	if err == nil {
		ticket.Issue()
	} else {
		fmt.Println(err)
	}
}

func initiateUnParkAndGenerateReceipt(commands []string) {
	ticketNumber := commands[len(commands)-1]
	ticketID, _ := strconv.Atoi(strings.TrimLeft(ticketNumber, "0"))
	receipt, err := parkingLot.UnPark(ticketID)
	if err == nil {
		receipt.Generate()
	} else {
		fmt.Println(err)
	}
}

func commandAllowed() []string {
	return []string{
		"park", "unpark",
	}
}

func contains(source []string, input string) bool {
	for _, a := range source {
		if strings.EqualFold(a, input) {
			return true
		}
	}
	return false
}

func vehiclesAllowed() []string {
	var result []string
	for _, vehicles := range spot.SlotVehicleMapping {
		result = append(result, vehicles...)
	}

	return result
}
