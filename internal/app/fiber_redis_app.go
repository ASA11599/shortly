package app

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type FiberRedisApp struct {
	fiberApp *fiber.App
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
	err := fra.initRedisClient()
	if err != nil { return err }
	return fra.initFiberApp()
}

func (fra *FiberRedisApp) Stop() error {
	err := fra.stopFiberApp()
	if err != nil { return err }
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

func (fra *FiberRedisApp) initFiberApp() error {
	fra.fiberApp = fiber.New()
	return fra.fiberApp.Listen("0.0.0.0:8080")
}

func (fra *FiberRedisApp) closeRedisClient() error {
	return fra.redisClient.Close()
}

func (fra *FiberRedisApp) stopFiberApp() error {
	return fra.fiberApp.Shutdown()
}
