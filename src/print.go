package main

import (
	"fmt"
	"strings"
)

func IsComplete(word string) {

}

func Display(word string, str string) {
	for _, letter := range word {
		if strings.Contains(string(letter), str) {
			fmt.Print(str)
		} else {
			fmt.Print("_")
		}
	}
}

func makeSlice(word string, input string) []string {
	var tab []string
	if IsInputValid(word, input) == true {
		tab = append(tab, input)
	}
	return tab
}

func IsInputValid(word string, input string) bool {
	if len(input) == 1 {
		for _, letter := range word {
			if string(letter) == input {
				return true
			}
		}
	} else {
		return input == word
	}
	return false
}
