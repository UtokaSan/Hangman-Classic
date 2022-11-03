package main

import (
	"fmt"
	"math/rand"
)

type HangmanData struct {
	Word            string
	RandomLetter    string
	Attempts        int
	HangmanPosition []string
}

func Random(nbr int) int {
	return rand.Intn(nbr)
}

func main() {
	var inputMenu int
	fmt.Println("---- [Menu] ----")
	fmt.Println("1. Play\n2. Rules")
	fmt.Print("Choose : ")
	fmt.Scanf("%d", &inputMenu)
	if inputMenu == 1 {
		game()
	}
	if inputMenu == 2 {
		rules()
	}
}

func rules() {
	fmt.Println("---- [Rules] ----")
	fmt.Print("Hangman is a simple word guessing game. Players try to figure out an unknown word by guessing letters. If too many letters which do not appear in the word are guessed, the player is hanged (and loses).")
}
