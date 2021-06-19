package basket

import (
	"context"

	"github.com/metinogurlu/go-basket/internal/adapters"
	"github.com/metinogurlu/go-basket/internal/app/commands"
	"github.com/metinogurlu/go-basket/internal/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(ctx context.Context) Application {
	repo := adapters.NewRedisRepository()

	return Application{
		Commands: Commands{
			UpdateBasket: commands.NewUpdateBasketHandler(repo),
			AddToBasket:  commands.NewAddToBasketHandler(repo),
		},
		Queries: Queries{
			GetBasket: queries.NewGetBasketHandler(repo),
		},
	}
}

type Queries struct {
	GetBasket queries.GetBasketHandler
}

type Commands struct {
	UpdateBasket commands.UpdateBasketHandler
	AddToBasket  commands.AddToBasketHandler
}
