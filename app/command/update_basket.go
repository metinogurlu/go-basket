package command

import (
	basket "basket/domain"
	"context"
)

type UpdateBasketHandler struct {
	repo basket.Repository
}

func NewUpdateBasketHandler(repo basket.Repository) UpdateBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return UpdateBasketHandler{repo}
}

func (h UpdateBasketHandler) Handle(ctx context.Context, id string, b basket.CustomerBasket) error {
	err := h.repo.UpdateBasket(ctx, id, b)

	if err != nil {
		return err
	}

	return nil
}
