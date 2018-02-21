package main

import (
	"log"
	"net/http"

	"github.com/sonnythehottest/restaurantapp/reservation"
	r "github.com/sonnythehottest/restaurantapp/restaurant"
)

func main() {
	port := "8080"

	m, err := r.NewModule(r.DefaultInitSample)
	if err != nil {
		log.Fatalln("failed to init restaurant module")
	}

	http.HandleFunc("/", m.HandleGetList())
	http.HandleFunc("/list", m.HandleGetList())
	http.HandleFunc("/restaurant", m.HandleGetInfo())
	http.HandleFunc("/reserve", reservation.HandleReserve(m))

	log.Println("listening to port", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
