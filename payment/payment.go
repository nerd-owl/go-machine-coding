package payment

import (
	"math"
	"time"

	"example.com/parking-lot/ticket"
)

type Payment interface {
	CalculateFee(t ticket.Ticket) float32
	GetRate() float32
}

type OfficeParkingPayment struct {
	Rate float32
}

func (o *OfficeParkingPayment) CalculateFee(t ticket.Ticket) float32 {
	timeInHours := math.Ceil(time.Since(t.GetCheckinTime()).Hours())
	return float32(timeInHours) * o.Rate
}

func (o *OfficeParkingPayment) GetRate() float32 {
	return o.Rate
}

func GetNewPayment(rate float32) Payment {
	return &OfficeParkingPayment{
		Rate: rate,
	}
}
