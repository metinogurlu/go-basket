package command

import (
	basket "basket/domain"
	"context"
)

type AddToBasketHandler struct {
	repo basket.Repository
}

func NewAddToBasketHandler(repo basket.Repository) AddToBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return AddToBasketHandler{repo}
}

func (h AddToBasketHandler) Handle(ctx context.Context, id string, item basket.BasketItem) error {
	err := h.repo.AddToBasket(ctx, id, item)

	if err != nil {
		return err
	}

	return nil
}
