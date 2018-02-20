package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sonnythehottest/restaurantapp/reservation"
	r "github.com/sonnythehottest/restaurantapp/restaurant"
)

func main() {

	m, err := r.NewModule()
	if err != nil {
		log.Fatalln("failed to init restaurant module")
	}

	http.HandleFunc("/", m.HandleGetList())
	http.HandleFunc("/list", m.HandleGetList())
	http.HandleFunc("/restaurant", m.HandleGetInfo())
	http.HandleFunc("/reserve", reservation.HandleReserve(m))

	port := "8080"
	fmt.Println("listening to port", port)
	http.ListenAndServe(":"+port, nil)
}
