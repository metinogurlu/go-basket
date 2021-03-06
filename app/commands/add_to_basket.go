package commands

import (
	"basket/config"
	models "basket/domain"
	"context"

	. "github.com/ahmetb/go-linq/v3"
)

type AddToBasketHandler struct {
	repo          models.Repository
	configuration config.Configuration
}

func NewAddToBasketHandler(repo models.Repository) AddToBasketHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return AddToBasketHandler{repo, config.NewConfiguration()}
}

func (h AddToBasketHandler) Handle(ctx context.Context, id string, itemToAdd models.BasketItem) error {
	b, err := h.repo.GetBasket(ctx, id)

	if err != nil {
		return err
	}

	if b.HasSameItem(itemToAdd) {
		i := From(*b.Items).IndexOfT(func(i models.BasketItem) bool {
			return i.Equals(itemToAdd)
		})
		basketItem := &(*b.Items)[i]

		if itemToAdd.Quantity < 0 {
			itemToAdd.Quantity = 0
		}
		// TODO: Handle basket rules
		if basketItem.Quantity+itemToAdd.Quantity <=
			h.configuration.BasketRules.MaximumSameItemInBasket {
			basketItem.Quantity += itemToAdd.Quantity
		}
	} else if len(*b.Items) < h.configuration.BasketRules.MaximumItemsInBasket {
		items := append(*b.Items, itemToAdd)
		b.Items = &items
	}

	err = h.repo.UpdateBasket(ctx, id, b)

	if err != nil {
		return err
	}

	return nil
}
