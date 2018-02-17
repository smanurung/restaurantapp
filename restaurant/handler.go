package restaurant

import (
	"net/http"
)

func HandleGetNearest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	return
}
