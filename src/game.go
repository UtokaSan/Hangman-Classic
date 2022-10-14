package main

import (
	"fmt"
)

func game() {
	var hangman HangmanData
	input := ""
	hangman.Word = Dictionnary("words/words.txt")
	fmt.Println(hangman.Word)
	str := string(hangman.Word[len(hangman.Word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	fmt.Scanf("%s", &input)
	test(hangman.Word, str, input)
}
