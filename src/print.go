package main

import (
	"fmt"
	"strings"
)

// Vérifie si il y a toujours des tirets
func VerifyDash(word string) bool {
	for _, el := range word {
		if strings.Contains(string(el), "_") {
			return false
		}
	}
	return true
}

// Print le len et les tirets
func test(word string, str string, input string) {
	input = strings.ToLower(input)
	for _, el := range word {
		if strings.Contains(string(el), str) {
			word = str
			fmt.Print(str)
		} else if strings.Contains(string(el), input) {
			word = input
			fmt.Print(input)
		} else {
			word = "_"
			fmt.Print(word)
		}
		if VerifyDash(word) == false {
			fmt.Scanf("%s", &input)
		}
	}
}
