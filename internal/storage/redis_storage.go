package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	redisClient *redis.Client
}

var redisStorage *RedisStorage

func GetRedisStorage() (*RedisStorage, error) {
	if redisStorage == nil {
		redisStorage = &RedisStorage{
			redisClient: nil,
		}
		return redisStorage, redisStorage.connect()
	}
	return redisStorage, nil
}

func (rs *RedisStorage) connect() error {
	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil { return err }
	rs.redisClient = redis.NewClient(options)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return rs.redisClient.Ping(ctx).Err()
}

func (rs *RedisStorage) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := rs.redisClient.Get(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (rs *RedisStorage) Set(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return rs.redisClient.Set(ctx, key, value, 0).Err()
}

func (rs *RedisStorage) Close() error {
	return rs.redisClient.Close()
}
