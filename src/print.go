package main

import "fmt"

func Printlen(word string, str string, input string) bool {
	for _, el := range word {
		transformString := string(el)
		if transformString == input {
			fmt.Print(input)
		} else if transformString == str {
			fmt.Print(str)
		} else {
			fmt.Print("_")
		}
	}
	return false
}
