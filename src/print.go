package main

import (
	"fmt"
	"strings"
)

// Vérifie si il y a toujours des tirets
func Printlen(word string) bool {
	for _, el := range word {
		if strings.Contains(string(el), "_") {
			return false
		}
	}
	return true
}

// Print le len et les tirets
func test(word string, str string, input string) {
	for _, el := range word {
		if strings.Contains(string(el), str) {
			fmt.Print(str)
		} else if strings.Contains(string(el), input) {
			fmt.Print(input)
		} else {
			fmt.Print("_")
		}
	}
	l := Printlen(word)
	fmt.Print(l)
}
