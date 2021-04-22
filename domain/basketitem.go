package basket

type BasketItem struct {
	ProductId  string
	Quantity   int
	Components *[]BasketItem
}
