package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func Random(nbr int) int {
	return rand.Intn(nbr)
}

// faire une boucle pour répéter l'input à chaque fois
func test(word string, str string, input string) {
	for Printlen(word, str, input) == true {
		fmt.Scanf("%s", &input)
	}
}

// function pour afficher
func Printlen(word string, str string, input string) bool {
	valid := false
	for i := 0; i < len(word); i++ {
		if string(word[i]) == str {
			fmt.Print(str)
		} else if string(word[i]) == input {
			fmt.Print(input)
		} else {
			fmt.Print("_")
			valid = true
		}
	}
	return valid
}

func game() {
	input := ""
	word := readFile("words/words.txt")
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	test(word, str, input)
}

func readFile(fileAccess string) string {
	data, _ := ioutil.ReadFile(fileAccess)
	str := string(data)
	split := strings.Split(str, "\n")
	random := Random(len(split))
	word := split[random]
	return word
}

func main() {
	game()
}
