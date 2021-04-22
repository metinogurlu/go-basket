package controller

import (
	models "basket/domain"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BasketController struct {
	repo models.Repository
}

func NewBasketController(repo models.Repository) BasketController {
	return BasketController{
		repo: repo,
	}
}

func (c *BasketController) GetBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	basket, _ := c.repo.GetBasket(params["id"])

	basketJson, _ := json.Marshal(basket)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(basketJson))
}

func (c *BasketController) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var basket models.CustomerBasket

	err := json.NewDecoder(r.Body).Decode(&basket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.repo.UpdateBasket(params["id"], &basket)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(params["id"]))
}
