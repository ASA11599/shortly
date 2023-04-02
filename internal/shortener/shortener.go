package shortener

import (
	"github.com/ASA11599/shortly/internal/alias"
	"github.com/ASA11599/shortly/internal/storage"
)

func Shorten(store storage.Storage, url string) (string, error) {
	a := alias.GenerateAlias()
	err := store.Set(a, url)
	return a, err
}

func Expand(store storage.Storage, alias string) (string, error) {
	return store.Get(alias)
}
