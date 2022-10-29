package main

func IsComplete(word string) {

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
