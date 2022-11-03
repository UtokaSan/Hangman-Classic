package main

import (
	"fmt"
)

func game() {
	var hangman HangmanData
	hangman.Word = Dictionnary("words/words.txt")
	hangman.RandomLetter = string(hangman.Word[len(hangman.Word)/2-1])
	for !IsHangmanComplete(hangman.Word, hangman.RandomLetter) {
		fmt.Print("\nChoose : \n")
		var input string
		fmt.Scanf("%s", &input)
		hangman.RandomLetter += input
		Display(hangman.Word, hangman.RandomLetter)
	}
	fmt.Print("\n\nCongrats !")
}
