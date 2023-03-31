package app

import (
	"context"
	"time"

	"github.com/ASA11599/shortly/internal/handlers"
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
	handlers.RegisterHandlers(fra.fiberApp)
	return fra.fiberApp.Listen("0.0.0.0:8080")
}

func (fra *FiberRedisApp) closeRedisClient() error {
	if fra.redisClient == nil { return nil }
	return fra.redisClient.Close()
}

func (fra *FiberRedisApp) stopFiberApp() error {
	if fra.fiberApp == nil { return nil }
	return fra.fiberApp.Shutdown()
}

func (fra *FiberRedisApp) Set(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return fra.redisClient.Set(ctx, key, value, 0).Err()
}

func (fra *FiberRedisApp) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := fra.redisClient.Get(ctx, key)
	return cmd.String(), cmd.Err()
}
