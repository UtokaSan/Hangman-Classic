package main

import (
	"math/rand"
)

type HangmanData struct {
	Word            string
	ToFind          string
	Attempts        int
	HangmanPosition []string
}

func Random(nbr int) int {
	return rand.Intn(nbr)
}

func main() {
	game()
}
