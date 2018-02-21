package reservation

// BasicReservation is a form of reservation of restaurant
type BasicReservation struct {
	ID                uint
	RestaurantID      uint
	Time              string
	TimexRestaurantID string
}

// NewBasicReservation returns BasicReservation struct with specific parameter
func NewBasicReservation(id, restaurantID uint, time, timexrestaurantID string) *BasicReservation {
	return &BasicReservation{id, restaurantID, time, timexrestaurantID}
}

// GetUniqueID returns unique identifier for reservation entry.
// In this case it's the string union of time (with hour precision) and restaurantID
func (br BasicReservation) GetUniqueID() string {
	return br.TimexRestaurantID
}

// GetData returns the struct for the reservation data
func (br BasicReservation) GetData() interface{} {
	return br
}
