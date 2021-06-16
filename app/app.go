package app

import (
	"basket/app/commands"
	"basket/app/query"
	"basket/infrastructure/adapters"
	"context"
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
			GetBasket: query.NewGetBasketHandler(repo),
		},
	}
}

type Queries struct {
	GetBasket query.GetBasketHandler
}

type Commands struct {
	UpdateBasket commands.UpdateBasketHandler
	AddToBasket  commands.AddToBasketHandler
}
