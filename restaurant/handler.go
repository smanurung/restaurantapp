package restaurant

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleGetList
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
			d = 5
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

// HandleGetInfo
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
