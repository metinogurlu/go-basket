package basket

type Repository interface {
	GetBasket(id string) (*CustomerBasket, error)
	UpdateBasket(id string, basket *CustomerBasket)
	AddToBasket(id string, basketItem *BasketItem) error
}
