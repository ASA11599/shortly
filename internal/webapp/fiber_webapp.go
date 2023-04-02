package webapp

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/ASA11599/shortly/internal/alias"
	"github.com/ASA11599/shortly/internal/shortener"
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
		a := c.Params("alias", "")
		if !alias.ValidateAlias(a) {
			return c.SendStatus(http.StatusBadRequest)
		}
		expanded, err := shortener.Expand(fwa.storage, a)
		if (err != nil) { return err }
		if _, err := url.ParseRequestURI(expanded); err != nil {
			return c.SendStatus(http.StatusNotFound)
		}
		return c.Redirect(expanded)
	})
	fwa.fiberApp.Post("/", func(c *fiber.Ctx) error {
		link := string(c.Body())
		if _, err := url.ParseRequestURI(link); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		a, err := shortener.Shorten(fwa.storage, link)
		if err != nil { return err }
		return c.SendString(a)
	})
}
