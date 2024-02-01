package ticket

import "time"

type Ticket interface {
	GetVehicleNo() string
	GetVehicleType() string
	GetCheckinTime() time.Time
	GetParkingSpot() string
}

type OfficeTicket struct {
	vehicleNo   string
	vehicleType string
	checkinTime time.Time
	parkingSpot string
}

func (o *OfficeTicket) GetVehicleNo() string {
	return o.vehicleNo
}

func (o *OfficeTicket) GetVehicleType() string {
	return o.vehicleType
}

func (o *OfficeTicket) GetParkingSpot() string {
	return o.parkingSpot
}

func (o *OfficeTicket) GetCheckinTime() time.Time {
	return o.checkinTime
}

func GetNewTicket(parkingSpace string, vehicleNo string, vehicleType string) Ticket {
	return &OfficeTicket{
		vehicleNo:   vehicleNo,
		vehicleType: vehicleType,
		checkinTime: time.Now(),
		parkingSpot: parkingSpace,
	}
}
