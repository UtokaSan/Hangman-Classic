package main

import (
	"fmt"
)

func Menu() {
	var answer int
	fmt.Println("Hangman\n")
	fmt.Println("----- Menu -----")
	fmt.Println("1. Play")
	fmt.Println("2. Rules")
	fmt.Println("----------------")
	fmt.Println("Choose a number")
	fmt.Scanf("%d", &answer)

	chooseDifficulty(answer)
}
func Rules() {
	fmt.Print()
}

func chooseDifficulty(answer int) {
	if answer == 1 {
		Game()
	} else if answer == 2 {
		Rules()
	} else {
		fmt.Print("Not possible")
	}
}

func Game() {
	word := "Hello"
	fmt.Println("Good Luck, you have 10 attempts")
	fmt.Println("Choose :")
	for ,el := range word {

	}
}

func main() {
	Menu()
}
