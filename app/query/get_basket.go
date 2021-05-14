package query

import (
	basket "basket/domain"
	"context"
)

type GetBasketHandler struct {
	repo basket.Repository
}

func NewGetBasketHandler(repo basket.Repository) GetBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetBasketHandler{repo}
}

func (h GetBasketHandler) Handle(ctx context.Context, id string) (*basket.CustomerBasket, error) {
	b, err := h.repo.GetBasket(ctx, id)
	if err != nil {
		return &basket.CustomerBasket{}, err
	}

	return b, nil
}
