package webapp

import (
	"errors"

	"github.com/ASA11599/shortly/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type FiberWebApp struct {
	fiberApp *fiber.App
	storage storage.Storage
}

var fiberWebApp *FiberWebApp

func GetFiberWebApp(storage storage.Storage) (*FiberWebApp, error) {
	if fiberWebApp == nil {
		fiberWebApp = &FiberWebApp{
			fiberApp: fiber.New(),
			storage: storage,
		}
		if storage == nil { return fiberWebApp, errors.New("FiberWebApp: Storage must be provided") }
		return fiberWebApp, nil
	} else {
		if storage != nil { return fiberWebApp, errors.New("FiberWebApp: Storage is already set") }
		return fiberWebApp, nil
	}
}

func (fwa *FiberWebApp) Start() error {
	fwa.registerHandlers()
	return fwa.fiberApp.Listen("0.0.0.0:8080")
}

func (fwa *FiberWebApp) Stop() error {
	if fwa.fiberApp == nil {
		return errors.New("the Fiber app was not initialized")
	}
	return fwa.fiberApp.Shutdown()
}

func (fwa *FiberWebApp) registerHandlers() {
	fwa.fiberApp.Get("/:alias", func(c *fiber.Ctx) error {
		alias := c.Params("alias", "")
		return c.SendString("redirecting " + alias)
	})
	fwa.fiberApp.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("creating alias for " + string(c.Body()))
	})
}
