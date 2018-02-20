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

// Module
//
// l is doubly linked list. The list is sorted by rating
type Module struct {
	l *list.List
	m map[int]*list.Element
}

// RatingType
type RatingType int

const (
	Bad RatingType = iota
	Fair
	Good
)

type R struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	CuisineType   string     `json:"cuisine_type"`
	Location      string     `json:"location"`
	Rating        RatingType `json:"rating"`
	Position      Point      `json:"position"`
	reservationDB *memdb.MemDB
}
