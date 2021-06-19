package commands

import (
	"context"

	"github.com/metinogurlu/go-basket/internal/adapters"
	"github.com/metinogurlu/go-basket/pkg/models"
)

type UpdateBasketHandler struct {
	repo adapters.Repository
}

func NewUpdateBasketHandler(repo adapters.Repository) UpdateBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return UpdateBasketHandler{repo}
}

func (h UpdateBasketHandler) Handle(ctx context.Context, id string, b models.CustomerBasket) error {
	err := h.repo.UpdateBasket(ctx, id, b)

	if err != nil {
		return err
	}

	return nil
}
