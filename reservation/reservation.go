package reservation

type BasicReservation struct {
	ID                uint
	RestaurantID      uint
	Time              string
	TimexRestaurantID string
}

func (br BasicReservation) GetUniqueID() string {
	return br.TimexRestaurantID
}

// Q: why I can't pass pointer as receiver here as interface method?
func (br BasicReservation) GetData() interface{} {
	return br
}
