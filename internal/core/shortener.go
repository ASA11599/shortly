package core

import (
	"math/rand"

	"github.com/ASA11599/shortly/internal/storage"
)

func Shorten(s storage.Store, url string) string {
	alias := randomString(5)
	s.Set(alias, url)
	return alias
}

func Expand(s storage.Store, alias string) string {
	return s.Get(alias)
}

func randomString(size int) string {
	res := make([]byte, size)
	for i := range res {
		res[i] = randomLetter()
	}
	return string(res)
}

func randomLetter() byte {
	return byte(97 + rand.Intn(25))
}
