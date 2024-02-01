package main

import (
	"fmt"

	"example.com/parking-lot/entry_gate"
	"example.com/parking-lot/exit_gate"
	"example.com/parking-lot/parking"
	"example.com/parking-lot/ticket"
)

func main() {
	var tickets []ticket.Ticket
	p := parking.GetNewParking([]string{"A1", "B1", "C1"})
	// pi := payment.GetNewPayment(10.00)
	entry := entry_gate.GetNewEntryGate(p)
	exit := exit_gate.GetNewExitGate(p)
	for i := 0; i < 4; i++ {
		t, err := entry.CreateTicket("UP70CN7150", "bike")
		if err != nil {
			fmt.Println(i+1, err.Error())
			continue
		}
		fmt.Println("ticket", t)
		tickets = append(tickets, t)
	}

	exit.ProcessVehicleExit(tickets[2])
	t, _ := entry.CreateTicket("UP70CN7150", "bike")
	fmt.Println(t)

	// for _, t := range tickets {
	// 	fmt.Printf("please pay %f$\n", pi.CalculateFee(t))
	// 	exit.ProcessVehicleExit(t)
	// }
}
