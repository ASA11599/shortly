package app

import (
	"errors"

	"github.com/ASA11599/shortly/internal/storage"
	"github.com/ASA11599/shortly/internal/webapp"
)

type ShortlyApp struct {
	storage storage.Storage
	webApp webapp.WebApp
}

var shortlyApp *ShortlyApp

func GetShortlyApp() *ShortlyApp {
	if shortlyApp == nil {
		return &ShortlyApp{
			storage: nil,
			webApp: nil,
		}
	}
	return shortlyApp
}

func (sa *ShortlyApp) Start() error {
	rs, err := storage.GetMemoryStorage()
	if err != nil { return err }
	sa.storage = rs
	sa.webApp, err = webapp.GetFiberWebApp(sa.storage)
	if err != nil { return err }
	sa.webApp.Start()
	return nil
}

func (sa *ShortlyApp) Stop() error {
	if (sa.webApp == nil) || (sa.storage == nil) {
		return errors.New("ShortlyApp could not start")
	}
	err := sa.webApp.Stop()
	if err != nil { return err }
	return sa.storage.Close()
}
