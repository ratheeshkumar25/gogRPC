package config

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisService represents the client of redis.
type RedisService struct {
	Client *redis.Client
}

// SetupRedis initializes redis server with configuration variables.
func SetupRedis() (*RedisService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_PORT"),
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("failed to connect to redis")
	}

	return &RedisService{
		Client: client,
	}, nil
}

// SetDataInRedis sets data in redis with a key and expiration time.
func (r *RedisService) SetDataInRedis(key string, value []byte, expTime time.Duration) error {
	err := r.Client.Set(context.Background(), key, value, expTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromRedis helps to retrieve the data from redis.

func (r *RedisService) GetFromRedis(key string) (string, error) {
	ctx := context.Background()
	jsonData, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return jsonData, nil
}

// func (r *RedisService) GetFromRedis(key string) (string, error) {
// 	jsonData, err := r.Client.Get(context.Background(), key).Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	return jsonData, nil
// }

// package config

// import (
// 	"context"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// )

// type RedisService struct {
// 	Client *redis.Client
// }

// func NewRedisService(addr string) *RedisService {
// 	client := redis.NewClient(&redis.Options{
// 		Addr: addr,
// 		DB:   0,
// 	})
// 	return &RedisService{Client: client}
// }

// func (r *RedisService) SetDataInRedis(key string, value []byte, expTime time.Duration) error {
// 	ctx := context.Background()
// 	err := r.Client.Set(ctx, key, value, expTime).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *RedisService) GetFromRedis(key string) (string, error) {
// 	ctx := context.Background()
// 	jsonData, err := r.Client.Get(ctx, key).Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	return jsonData, nil
// }
