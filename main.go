package main

import (
	"basket/app"
	controller "basket/interfaces"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	ctx := context.Background()
	app := app.NewApplication(ctx)
	c := controller.NewBasketController(app)

	r.HandleFunc("/baskets/{id}", c.GetBasket).Methods("GET")
	r.HandleFunc("/baskets/{id}", c.UpdateBasket).Methods("PUT")
	r.HandleFunc("/baskets/{id}", c.AddtoBasket).Methods("POST")

	port := ":9000" // port for run the app

	fmt.Println("Start listening on port", port) // server up

	// Serve server on a port
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err) // print error log
	}

}
