package main

import (
	"fmt"
)

func game() {
	var hangman HangmanData
	hangman.Word = Dictionnary("words/words.txt")
	fmt.Println(hangman.Word)
	hangman.Str = string(hangman.Word[len(hangman.Word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	input := ""
	fmt.Scanf("%s", &input)
	makeSlice(hangman.Word, input)
	Display(hangman.Word, hangman.Str)
}
