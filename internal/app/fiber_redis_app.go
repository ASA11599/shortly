package app

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type FiberRedisApp struct {
	redisClient *redis.Client
}

var fiberRedisApp *FiberRedisApp

func GetInstance() *FiberRedisApp {
	if fiberRedisApp == nil { 
		fiberRedisApp =  &FiberRedisApp{
			redisClient: nil,
		}
	 }
	return fiberRedisApp
}

func (fra *FiberRedisApp) Start() error {
	return fra.initRedisClient()
}

func (fra *FiberRedisApp) Stop() error {
	return fra.closeRedisClient()
}

func (fra *FiberRedisApp) initRedisClient() error {
	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil { return err }
	fra.redisClient = redis.NewClient(options)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return fra.redisClient.Ping(ctx).Err()
}

func (fra *FiberRedisApp) closeRedisClient() error {
	return fra.redisClient.Close()
}
