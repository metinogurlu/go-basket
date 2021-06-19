package adapters

import (
	"context"

	"github.com/metinogurlu/go-basket/pkg/models"
)

type Repository interface {
	GetBasket(ctx context.Context, id string) (models.CustomerBasket, error)
	UpdateBasket(ctx context.Context, id string, basket models.CustomerBasket) error
}
