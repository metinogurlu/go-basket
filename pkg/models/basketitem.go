package models

import . "github.com/ahmetb/go-linq/v3"

type BasketItem struct {
	ProductId  string
	Quantity   int
	Components []BasketItem
}

func (i BasketItem) Equals(other BasketItem) bool {
	return i.ProductId == other.ProductId && i.componentsEquals(other)
}

func (i BasketItem) componentsEquals(other BasketItem) bool {
	if (i.Components == nil || len(i.Components) == 0) &&
		(other.Components == nil || len(other.Components) == 0) {
		return true
	}

	if len(i.Components) != len(other.Components) {
		return false
	}

	for _, component := range i.Components {
		otherHasThis := From(other.Components).AnyWith(func(c interface{}) bool {
			return c.(BasketItem).ProductId == component.ProductId &&
				c.(BasketItem).Quantity == component.Quantity
		})

		if !otherHasThis {
			return false
		}

	}

	return true
}
