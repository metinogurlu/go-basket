package app

import (
	"basket/app/command"
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
			UpdateBasket: command.NewUpdateBasketHandler(repo),
			AddToBasket:  command.NewAddToBasketHandler(repo),
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
	UpdateBasket command.UpdateBasketHandler
	AddToBasket  command.AddToBasketHandler
}
