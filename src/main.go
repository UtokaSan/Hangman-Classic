package main

import (
	"math/rand"
	"time"
)

type HangmanData struct {
	Word            string
	ToFind          string
	Attempts        int
	HangmanPosition []string
}

func Random(nbr int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(nbr)
}

func main() {
	game()
}
