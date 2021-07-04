package portdomain

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

type Store interface {
	Add(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (value string, err error)
}

type portStoreRedis struct {
	client *redis.Client
}

func (ps *portStoreRedis) Add(ctx context.Context, key, value string) error {
	return ps.client.Set(ctx, key, value, 0).Err()
}

func (ps *portStoreRedis) Get(ctx context.Context, key string) (value string, err error) {
	return ps.client.Get(ctx, key).Result()
}

func NewPortStoreRedis() Store {
	return &portStoreRedis{
		client: redis.NewClient(&redis.Options{
			DB:   0,
			Addr: os.Getenv("REDIS"),
		}),
	}
}
