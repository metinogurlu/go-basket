package command

import (
	basket "basket/domain"
	"context"
	"math"

	. "github.com/ahmetb/go-linq/v3"
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
	b, err := h.repo.GetBasket(ctx, id)

	if err != nil {
		return err
	}

	if b.HasSameItem(item) {
		i := (From(*b.Items).FirstWith(func(c interface{}) bool {
			return c.(basket.BasketItem).Equals(item)
		}).(basket.BasketItem))

		quantityToAdd := int(math.Max(float64(item.Quantity), 1))
		quantityAfterAdd := item.Quantity + quantityToAdd

		// TODO: check quantity referance
		if quantityAfterAdd <= 20 {
			i.Quantity = quantityAfterAdd
		}
	} else if len(*b.Items) < 50 {
		*b.Items = append(*b.Items, item)
		h.repo.UpdateBasket(ctx, id, b)
	}

	return nil
}
