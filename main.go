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
	feeFactory = feemodel.NewFeeFactory()
	parkingLot = parkinglot.NewParkingLot(
		map[enum.SpotType]int{
			enum.TwoWheelers:      3,
			enum.SmallFourWheeler: 3,
		}, enum.Mall, clock.NewClock(), feeFactory)
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

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

func initiateParkAndIssueTicket(commands []string) {
	ticket := parkingLot.Park(commands[1])
	if ticket != nil {
		ticket.Issue()
	} else {
		fmt.Printf("The parking is full for %s. Please come back later. \n", commands[1])
	}
}

func initiateUnParkAndGenerateReceipt(commands []string) {
	ticketNumber := commands[len(commands)-1]
	fmt.Println(commands)
	ticketID, _ := strconv.Atoi(ticketNumber)
	receipt := parkingLot.UnPark(ticketID)
	if receipt != nil {
		receipt.Generate()
	} else {
		fmt.Println("Please enter correct ticket id without the leading zeros")
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
