package main

import (
	"fmt"
)

func isEnoughChangeForTickets(moneyForTickets []int) string {

	currentChange := 0
	const ticketCost = 25
	for _, moneyForTicket := range moneyForTickets {
		if moneyForTicket - ticketCost
	}

	return "NO"
}

func doTestIsEnoughChangeForTickets(moneyForTickets []int, expectedResult bool) {
	isEnoughChangeForTicketsResult := isEnoughChangeForTickets(moneyForTickets)
	expectedStringResult := "NO"
	if true == expectedResult {
		expectedStringResult = "YES"
	}
	if isEnoughChangeForTicketsResult != expectedStringResult {
		fmt.Printf("TEST NOT PASSED. moneyForTickets: %v, expectedResult: %s, isEnoughChangeForTicketsResult: %s\n", moneyForTickets, expectedStringResult, isEnoughChangeForTicketsResult)

		return
	}

	fmt.Printf("test passed. moneyForTickets: %v, expectedResult: %s, isEnoughChangeForTicketsResult: %s\n", moneyForTickets, expectedStringResult, isEnoughChangeForTicketsResult)
}

func main() {
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50},
		true,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 100},
		false,
	)
	doTestIsEnoughChangeForTickets(
		[]int{25, 25, 50, 50, 100},
		false,
	)
}