package entities

import (
	"math/rand"
	"time"
)

func GetNames(quantity int) []string {
	var randInts []int
	var names []string

	max := len(namesData)

	for len(names) < quantity {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(max - 1)

		// Check if it already has the number
		exists := false
		for _, a := range randInts {
			if a == n {
				exists = true
				break
			}
		}

		if !exists {
			names = append(names, namesData[n])
		}
	}

	return names
}
