package main

import (
	"fmt"
	"net/http"

	r "github.com/sonnythehottest/restaurantapp/restaurant"
)

func main() {
	http.HandleFunc("/nearest", r.HandleGetNearest)

	port := "8080"
	fmt.Println("listening to port", port)
	http.ListenAndServe(":"+port, nil)
}
