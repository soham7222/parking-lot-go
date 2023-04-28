package main

import (
	"bufio"
	"fmt"
	"os"
	"sahaj-parking-lot/clock"
	"sahaj-parking-lot/enum"
	"sahaj-parking-lot/feemodel"
	"sahaj-parking-lot/parkinglot"
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

		if strings.Contains(input, "unpark") {
			initiateUnParkAndGenerateReceipt(input)
		} else if strings.Contains(input, "park") {
			initiateParkAndIssueTicket(input)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func initiateParkAndIssueTicket(input string) {
	words := strings.Split(input, " ")
	ticket := parkingLot.Park(words[len(words)-1])
	if ticket != nil {
		ticket.Issue()
	} else {
		fmt.Printf("The parking is full for %s. Please come back later. \n", words[len(words)-1])
	}
}

func initiateUnParkAndGenerateReceipt(input string) {
	words := strings.Split(input, " ")
	ticketNumber := words[len(words)-1]
	ticketID, _ := strconv.Atoi(ticketNumber)
	receipt := parkingLot.UnPark(ticketID)
	if receipt != nil {
		receipt.Generate()
	} else {
		fmt.Println("Please enter correct ticket id without the leading zeros")
	}
}
