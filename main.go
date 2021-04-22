package main

import (
	controller "basket/app"
	"basket/infrastructure/adapters"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var r = mux.NewRouter()
	var repo = adapters.NewRedisRepository()
	c := controller.NewBasketController(repo)

	r.HandleFunc("/baskets/{id}", c.GetBasket).Methods("GET")
	r.HandleFunc("/baskets/{id}", c.UpdateBasket).Methods("PUT")

	port := ":9000" // port for run the app

	fmt.Println("Start listening on port", port) // server up

	// Serve server on a port
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err) // print error log
	}

}
