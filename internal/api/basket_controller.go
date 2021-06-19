package api

import (
	"encoding/json"
	"net/http"

	basket "github.com/metinogurlu/go-basket/internal/app"
	"github.com/metinogurlu/go-basket/pkg/models"

	"github.com/gorilla/mux"
)

type BasketController struct {
	app basket.Application
}

func NewBasketController(app basket.Application) BasketController {
	return BasketController{app}
}

func (c BasketController) GetBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	b, _ := c.app.Queries.GetBasket.Handle(r.Context(), params["id"])

	basketJson, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(basketJson))
}

func (c BasketController) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var b models.CustomerBasket

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.app.Commands.UpdateBasket.Handle(r.Context(), params["id"], b)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(params["id"]))
}

func (c BasketController) AddtoBasket(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var item models.BasketItem

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.app.Commands.AddToBasket.Handle(r.Context(), id, item)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(item.ProductId))
}
