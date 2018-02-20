package restaurant

import (
	"log"
	"reflect"
	"testing"
)

func TestNewModule(t *testing.T) {
	m, err := NewModule()
	if err != nil {
		t.Fatalf("failed to create new module")
	}
	if !reflect.DeepEqual(m, md) {
		t.Errorf("modules are not deeply equal")
	}
}

func TestGetWithDistance(t *testing.T) {
	cases := []struct {
		from      Point
		distance  int
		lenResult int
	}{
		{Point{0, 0}, 2, 1},
		{Point{0, 0}, 1, 0},
		{Point{0, 0}, 15, 10},
	}

	for _, c := range cases {
		r := md.GetWithDistance(c.from, c.distance)
		if len(r) != c.lenResult {
			t.Errorf("md.GetWithDistance(): len %d, want %d", len(r), c.lenResult)
		}
	}
}

var (
	md  *Module
	err error
)

func TestMain(m *testing.M) {
	md, err = NewModule()
	if err != nil {
		log.Fatal("can't initialize new module")
	}

	m.Run()
}
