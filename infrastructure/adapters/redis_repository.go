package adapters

import (
	models "basket/domain"
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	db  *redis.Client
	ctx context.Context
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		db:  getRedisClient(),
		ctx: context.Background(),
	}
}

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (repo *RedisRepository) GetBasket(ctx context.Context, id string) (models.CustomerBasket, error) {
	var basket models.CustomerBasket

	s, _ := repo.db.Get(repo.ctx, id).Result()
	json.Unmarshal([]byte(s), &basket)

	return basket, nil
}

func (repo *RedisRepository) UpdateBasket(ctx context.Context, id string, b models.CustomerBasket) error {
	err := repo.db.Set(repo.ctx, id, b, 0)
	if err != nil {
		return err.Err()
	}

	return nil
}
