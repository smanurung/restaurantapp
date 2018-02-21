package restaurant

import (
	"container/list"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/go-memdb"
)

func TestGetWithDistance(t *testing.T) {
	cases := []struct {
		from     Point
		distance int
		result   []R
	}{
		{Point{0, 0}, 2, []R{
			R{0, "TestName", "TestCuisineType", "TestLocation", 0, Point{0, 0}, md.memDB},
			R{1, "TestName", "TestCuisineType", "TestLocation", 0, Point{1, 1}, md.memDB},
		}},
		{Point{0, 0}, 1, []R{
			R{0, "TestName", "TestCuisineType", "TestLocation", 0, Point{0, 0}, md.memDB},
		}},
		{Point{0, 0}, 15, []R{
			R{0, "TestName", "TestCuisineType", "TestLocation", 0, Point{0, 0}, md.memDB},
			R{1, "TestName", "TestCuisineType", "TestLocation", 0, Point{1, 1}, md.memDB},
			R{2, "TestName", "TestCuisineType", "TestLocation", 0, Point{2, 2}, md.memDB},
			R{3, "TestName", "TestCuisineType", "TestLocation", 0, Point{3, 3}, md.memDB},
			R{4, "TestName", "TestCuisineType", "TestLocation", 0, Point{4, 4}, md.memDB},
			R{5, "TestName", "TestCuisineType", "TestLocation", 0, Point{5, 5}, md.memDB},
			R{6, "TestName", "TestCuisineType", "TestLocation", 0, Point{6, 6}, md.memDB},
			R{7, "TestName", "TestCuisineType", "TestLocation", 0, Point{7, 7}, md.memDB},
			R{8, "TestName", "TestCuisineType", "TestLocation", 0, Point{8, 8}, md.memDB},
			R{9, "TestName", "TestCuisineType", "TestLocation", 0, Point{9, 9}, md.memDB},
			R{10, "TestName", "TestCuisineType", "TestLocation", 0, Point{10, 10}, md.memDB},
		}},
	}

	for _, c := range cases {
		r := md.GetWithDistance(c.from, c.distance)
		if !reflect.DeepEqual(r, c.result) {
			t.Errorf("md.GetWithDistance(): %v, want %v", r, c.result)
		}
	}
}

func TestGetRestaurant(t *testing.T) {
	cases := []struct {
		i      int
		result *R
	}{
		{0, &R{0, "TestName", "TestCuisineType", "TestLocation", 0, Point{0, 0}, md.memDB}},
		{1, &R{1, "TestName", "TestCuisineType", "TestLocation", 0, Point{1, 1}, md.memDB}},
	}

	for _, c := range cases {
		actual := md.GetRestaurant(c.i)
		if !reflect.DeepEqual(actual, c.result) {
			t.Errorf("md.GetRestaurant(): %v, want %v", actual, c.result)
		}
	}
}

func TestReserve(t *testing.T) {
	cases := []struct {
		r  *R
		rv Reservation
	}{
		{&R{0, "TestName", "TestCuisineType", "TestLocation", 0, Point{0, 0}, md.memDB}, DummyReservation{0, "20180221T130"}},
	}

	for _, c := range cases {
		if err := c.r.Reserve(c.rv); err != nil {
			t.Errorf("c.r.Reserve(): %v, want err is nil\n", err)
		}
		if err := c.r.Reserve(c.rv); err == nil {
			t.Errorf("c.r.Reserve(): %v, want err is not nil\n", err)
		}
	}
}

type DummyReservation struct {
	ID                uint
	TimexRestaurantID string
}

func (dr DummyReservation) GetUniqueID() string {
	return dr.TimexRestaurantID
}

func (dr DummyReservation) GetData() interface{} {
	return dr
}

func DummyInitSample(mdb *memdb.MemDB, num int) (*list.List, map[int]*list.Element) {
	l := list.New()
	m := make(map[int]*list.Element)

	for i := 0; i < num; i++ {
		r := R{
			ID:            i,
			Name:          "TestName",
			CuisineType:   "TestCuisineType",
			Location:      "TestLocation",
			Rating:        0,
			Position:      Point{i, i},
			reservationDB: mdb,
		}

		e := l.PushBack(r)
		m[r.ID] = e
	}
	return l, m
}

var (
	md  *Module
	err error
)

func TestMain(m *testing.M) {
	md, err = NewModule(DummyInitSample)
	if err != nil {
		log.Fatal("can't initialize new module")
	}

	os.Exit(m.Run())
}
