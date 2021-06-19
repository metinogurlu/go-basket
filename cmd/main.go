package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/metinogurlu/go-basket/internal/api"

	basket "github.com/metinogurlu/go-basket/internal/app"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	ctx := context.Background()
	app := basket.NewApplication(ctx)
	c := api.NewBasketController(app)

	r.HandleFunc("/baskets/{id}", c.GetBasket).Methods("GET")
	r.HandleFunc("/baskets/{id}", c.UpdateBasket).Methods("PUT")
	r.HandleFunc("/baskets/{id}", c.AddtoBasket).Methods("POST")

	port := ":9001" // port for run the app

	fmt.Println("Start listening on port", port) // server up

	// Serve server on a port
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err) // print error log
	}

}
