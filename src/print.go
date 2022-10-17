package main

import (
	"fmt"
	"strings"
)

// Vérifie si il y a toujours des tirets
func VerifyDash(word string, input string) bool {
	found := true
	for _, el := range word {
		if strings.Contains(string(el), "_") {
			found = false
		}
	}
	if found == false {
		fmt.Scanf("%s", &input)
	}
	return found
}

// Print le len et les tirets
func display(word string, str string, input string) string {
	for _, el := range word {
		if strings.Contains(string(el), str) {
			fmt.Print(string(el))
		} else {
			fmt.Print("_")
		}
	}
	RevealLetter(word, input)
	fmt.Print(VerifyDash(word, input))
	return word
}

func RevealLetter(word string, input string) {
	for i := 0; i < len(word); i++ {
		if string(word[i]) == input {

		}
	}
}
