package restaurant

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hashicorp/go-memdb"
)

// InitSampleFunc is a function type to generate restaurants sample structure.
// This returns doubly-linked list and map of restaurantID and element containing restaurant value.
// The doubly-linked list is sorted by rating in this case.
type InitSampleFunc func(*memdb.MemDB, int) (*list.List, map[int]*list.Element)

// NewModule returns restaurant module given sampling function for restaurant generation.
func NewModule(isf InitSampleFunc) (*Module, error) {

	// memDB is used to store the relation for reservation data.
	// DB is indexed by reservationID and composition of restaurantID and time (with hour precision in this example)
	// reservationID index is to support query for reservation detail retrieval
	// composition index is to support reserve operation, which is to check if certain timeslot for certain restaurant is already booked or not.
	dbSchema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"reservation": &memdb.TableSchema{
				Name: "reservation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "ID"},
					},
					"TimexRestaurantID": &memdb.IndexSchema{
						Name:    "TimexRestaurantID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "TimexRestaurantID"},
					},
				},
			},
		},
	}

	memDB, err := memdb.NewMemDB(dbSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to create memdb")
	}

	l, m := isf(memDB, 300) // initialize 300 sample restaurants

	return &Module{
		l:     l,
		m:     m,
		memDB: memDB,
	}, nil
}

// DefaultInitSample is default init sampling function to generate restaurant sample for this program.
// Values are mostly pseudo-random, e.g. restaurantID, cuisineType, rating, etc.
func DefaultInitSample(mdb *memdb.MemDB, num int) (*list.List, map[int]*list.Element) {

	l := list.New()
	m := make(map[int]*list.Element)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		r := R{
			ID:   rand.Intn(num),
			Name: fmt.Sprintf("restaurant sample #%d", i),
			CuisineType: map[int]string{
				0: "Asian",
				1: "Western",
				2: "Indonesian",
				3: "Italian",
			}[rand.Intn(4)],
			Location: fmt.Sprintf("address sample #%d", i),
			Rating: map[int]RatingType{
				0: Good,
				1: Fair,
				2: Bad,
			}[i/100],
			Position:      Point{i, i},
			reservationDB: mdb,
		}

		e := l.PushBack(r)
		m[r.ID] = e
	}
	return l, m
}

// GetWithDistance compares the distance between `from` and all points in restaurant list,
// find with distance less than or equal to `d`, sorted by restaurant rating (as the linked list is already sorted by rating)
func (m *Module) GetWithDistance(from Point, d int) []R {
	sqr := int(math.Pow(float64(d), 2))

	var rs []R

	for e := m.l.Front(); e != nil; e = e.Next() {
		r, ok := e.Value.(R)
		if !ok {
			continue
		}
		if cur := int(math.Pow(float64(r.Position.Lat-from.Lat), 2)) + int(math.Pow(float64(r.Position.Long-from.Long), 2)); cur <= sqr {
			rs = append(rs, r)
		}
	}

	return rs
}

// GetRestaurant return restaurant struct given restaurantID, return nil if not found
func (m *Module) GetRestaurant(id int) *R {
	itv, ok := m.m[id]
	if !ok {
		return nil
	}
	rs, ok := itv.Value.(R)
	if !ok {
		return nil
	}
	return &rs
}

// Reservation contains behavior for reservation struct
//
// GetUniqueID() returns string that represents reservation uniquely, in this case in terms of time & restaurantID (for next reservation check)
// GetData() returns object of the reservation itself, e.g. for further to be saved to DB
type Reservation interface {
	GetUniqueID() string
	GetData() interface{}
}

// Reserve add reservation entry to the db.
// This checks if reservation has been made before the particular restaurant at certain time (hour precision).
func (r *R) Reserve(rv Reservation) error {
	tx := r.reservationDB.Txn(true)

	curr, err := tx.First("reservation", "TimexRestaurantID", rv.GetUniqueID())
	if err != nil {
		tx.Abort()
		return err
	}

	if curr != nil {
		log.Printf("reservation has already been made before: %v\n", curr)
		tx.Abort()
		return fmt.Errorf("can't reserve, booked already")
	}

	err = tx.Insert("reservation", rv.GetData())
	if err != nil {
		tx.Abort()
		return err
	}

	tx.Commit()
	return nil
}
