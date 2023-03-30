package app

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type FiberRedisApp struct {
	redisClient *redis.Client
}

func NewFiberRedisApp() *FiberRedisApp {
	return &FiberRedisApp{
		redisClient: nil,
	}
}

func (fra *FiberRedisApp) Start() error {
	return fra.initRedisClient()
}

func (fra *FiberRedisApp) Stop() error {
	fmt.Println("Stopping app")
	return fra.closeRedisClient()
}

func (fra *FiberRedisApp) initRedisClient() error {
	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil { return err }
	fra.redisClient = redis.NewClient(options)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(fra.redisClient)
	time.Sleep(5 * time.Second)
	return fra.redisClient.Ping(ctx).Err()
}

func (fra *FiberRedisApp) closeRedisClient() error {
	return fra.redisClient.Close()
}
