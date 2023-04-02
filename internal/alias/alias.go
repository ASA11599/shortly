package alias

import (
	"math/rand"
)

func randomLetter() byte {
	return byte(97 + rand.Intn(25))
}

func GenerateAlias() string {
	// Generates a random string of 5 characters
	return string([]byte{
		randomLetter(),
		randomLetter(),
		randomLetter(),
		randomLetter(),
		randomLetter(),
	})
}

func ValidateAlias(alias string) bool {
	if len(alias) != 5 { return false }
	for _, b := range alias {
		if (b < 97) || (b > 122) { return false }
	}
	return true
}
