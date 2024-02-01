package entry_gate

import (
	"errors"

	"example.com/parking-lot/parking"
	"example.com/parking-lot/ticket"
)

type EntryGate interface {
	CreateTicket(vehicleNo string, vehicleType string) (ticket.Ticket, error)
}

type OfficeEntryGate struct {
	parking.Parking
}

func (o *OfficeEntryGate) CreateTicket(vehicleNo string,
	vehicleType string) (ticket.Ticket, error) {
	if o.IsSpaceLeft() {
		parkingSpace := o.Book()
		return ticket.GetNewTicket(parkingSpace, vehicleNo, vehicleType), nil
	}

	return &ticket.OfficeTicket{}, errors.New("no space available")
}

func GetNewEntryGate(p parking.Parking) EntryGate {
	return &OfficeEntryGate{
		p,
	}
}
