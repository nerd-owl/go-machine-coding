package exit_gate

import (
	"example.com/parking-lot/parking"
	"example.com/parking-lot/ticket"
)

type ExitGate interface {
	ProcessVehicleExit(t ticket.Ticket)
}

type OfficeExitGate struct {
	parking.Parking
}

func (o *OfficeExitGate) ProcessVehicleExit(t ticket.Ticket) {
	o.Unbook(t)
}

func GetNewExitGate(p parking.Parking) ExitGate {
	return &OfficeExitGate{
		p,
	}
}
