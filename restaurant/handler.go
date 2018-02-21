package restaurant

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleGetList handles request to list all restaurants within certain distance from the client.
// Restaurants returned are sorted by rating.
func (m *Module) HandleGetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		flat, err := strconv.Atoi(r.FormValue("lat"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		flong, err := strconv.Atoi(r.FormValue("long"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		d, err := strconv.Atoi(r.FormValue("distance"))
		if err != nil {
			d = 5 // default value for distance here is 5 (km)
		}

		rs := m.GetWithDistance(Point{flat, flong}, d)

		encoded, err := json.Marshal(rs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(encoded)
		return
	}
}

// HandleGetInfo handles request to list information on specific restaurant given restaurantID
func (m *Module) HandleGetInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rs := m.GetRestaurant(id)
		if rs == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		e, err := json.Marshal(rs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(e)
		return
	}
}
