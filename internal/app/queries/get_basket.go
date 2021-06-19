package queries

import (
	"context"

	"github.com/metinogurlu/go-basket/internal/adapters"
	"github.com/metinogurlu/go-basket/pkg/models"
)

type GetBasketHandler struct {
	repo adapters.Repository
}

func NewGetBasketHandler(repo adapters.Repository) GetBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return GetBasketHandler{repo}
}

func (h GetBasketHandler) Handle(ctx context.Context, id string) (models.CustomerBasket, error) {
	b, err := h.repo.GetBasket(ctx, id)
	if err != nil {
		return models.CustomerBasket{}, err
	}

	return b, nil
}
