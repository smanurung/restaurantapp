package reservation

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sonnythehottest/restaurantapp/restaurant"
)

// HandleReserve receives http request from client to reserve place at a restaurant,
// parameters are:
//   - restaurantID
//   - date (format: yyyymmddThh)
// In this version, it's assumed that a restaurant can only have 1 place at each hour.
//
// Example request:
//   curl '<host>:<port>/reserve?restaurantID=1&time=20180201T02'
func HandleReserve(m *restaurant.Module) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rsid, err := strconv.Atoi(r.FormValue("restaurantID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t := r.FormValue("time")
		if t == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rs := m.GetRestaurant(rsid)
		if rs == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = rs.Reserve(NewBasicReservation(uint(time.Now().UnixNano()), uint(rsid), t, fmt.Sprintf("%s%d", t, rsid)))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
