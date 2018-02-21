package restaurant

import (
	"container/list"

	"github.com/hashicorp/go-memdb"
)

// Point represents position on the map
// Lat & Long here will represent dimensions x & y axis
type Point struct {
	Lat  int `json:"latitude"`
	Long int `json:"longitude"`
}

// Module is restaurant module containing several structures for restaurant representation.
//
// l is doubly-linked list. The list is sorted by rating.
// m is map from restaurantID to list element, containing restaurant details.
// memDB is memory data store to save reservation data.
type Module struct {
	l     *list.List
	m     map[int]*list.Element
	memDB *memdb.MemDB
}

// RatingType is type for rating
type RatingType int

const (
	// Bad represents bad rating for restaurant
	Bad RatingType = iota
	// Fair represents fair rating for restaurant
	Fair
	// Good represents good rating for restaurant
	Good
)

// R is restaurant struct containing restaurant detail fields
type R struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	CuisineType   string     `json:"cuisine_type"`
	Location      string     `json:"location"`
	Rating        RatingType `json:"rating"`
	Position      Point      `json:"position"`
	reservationDB *memdb.MemDB
}
