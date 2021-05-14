package basket

import "context"

type Repository interface {
	GetBasket(ctx context.Context, id string) (CustomerBasket, error)
	UpdateBasket(ctx context.Context, id string, basket CustomerBasket) error
}
