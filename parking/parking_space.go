package parking

import "example.com/parking-lot/ticket"

type Parking interface {
	IsSpaceLeft() bool
	Book() string
	Unbook(t ticket.Ticket)
}

type OfficeParking struct {
	spacesMap map[string]bool
}

func (o *OfficeParking) IsSpaceLeft() bool {
	for _, v := range o.spacesMap {
		if !v {
			return true
		}
	}

	return false
}

func (o *OfficeParking) Book() string {
	var space string
	for k, v := range o.spacesMap {
		if !v {
			space = k
			o.spacesMap[k] = true
			return space
		}
	}

	return ""
}

func (o *OfficeParking) Unbook(t ticket.Ticket) {
	o.spacesMap[t.GetParkingSpot()] = false
}

func GetNewParking(spaces []string) Parking {
	mp := make(map[string]bool)

	for _, v := range spaces {
		mp[v] = false
	}

	return &OfficeParking{
		spacesMap: mp,
	}
}
