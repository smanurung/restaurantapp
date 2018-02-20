package restaurant

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hashicorp/go-memdb"
)

// NewModule
func NewModule() (*Module, error) {

	// initialize pre-defined restaurants data for the program

	rand.Seed(time.Now().UnixNano())
	l := list.New()
	m := make(map[int]*list.Element)

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

	n := 300 // number of restaurant samples
	for i := 0; i < n; i++ {
		r := R{
			ID:   rand.Intn(n),
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
			reservationDB: memDB,
		}

		e := l.PushBack(r)
		m[r.ID] = e
	}

	return &Module{
		l: l,
		m: m,
	}, nil
}

// GetWithDistance compares the distance between `from` and all points in restaurant list,
// find with distance lesser than `d`, sorted by restaurant rating
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

// GetRestaurant
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

type Reservation interface {
	GetUniqueID() string
	GetData() interface{}
}

// Reserve add reservation entry to the db
func (r *R) Reserve(rv Reservation) error {
	tx := r.reservationDB.Txn(true)

	existingRv, err := tx.First("reservation", "TimexRestaurantID", rv.GetUniqueID())
	if err != nil {
		tx.Abort()
		return err
	}

	if existingRv != nil {
		fmt.Printf("existing reservation %v\n", existingRv)
		tx.Abort()
		return fmt.Errorf("it's reserved already")
	}

	err = tx.Insert("reservation", rv.GetData())
	if err != nil {
		tx.Abort()
		return err
	}

	tx.Commit()

	return nil
}
