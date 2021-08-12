package entities

import (
	"github.com/brianvoe/gofakeit/v6"
)

func GetNames(quantity int) (names []string) {
	for len(names) < quantity {
		name := gofakeit.Name()

		// Check if it already has the number
		exists := false
		for _, a := range names {
			if a == name {
				exists = true
				break
			}
		}

		if !exists {
			names = append(names, name)
		}
	}

	return names
}
