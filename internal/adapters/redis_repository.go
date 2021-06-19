package adapters

import (
	"context"
	"encoding/json"

	"github.com/metinogurlu/go-basket/configs"
	"github.com/metinogurlu/go-basket/pkg/models"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	db            *redis.Client
	ctx           context.Context
	configuration configs.Configuration
}

func NewRedisRepository() *RedisRepository {
	c := configs.NewConfiguration()

	return &RedisRepository{
		db:            getRedisClient(c.Redis.ConnectionString),
		ctx:           context.Background(),
		configuration: c,
	}
}

func getRedisClient(connectionString string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (repo *RedisRepository) GetBasket(ctx context.Context, id string) (models.CustomerBasket, error) {
	var basket models.CustomerBasket

	s, err := repo.db.Get(repo.ctx, id).Result()
	json.Unmarshal([]byte(s), &basket)

	return basket, err
}

func (repo *RedisRepository) UpdateBasket(ctx context.Context, id string, b models.CustomerBasket) error {
	err := repo.db.Set(repo.ctx, id, b, 0)
	if err != nil {
		return err.Err()
	}

	return nil
}
