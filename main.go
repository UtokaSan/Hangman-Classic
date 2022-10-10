package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func Random(nbr int) int {
	return rand.Intn(nbr)
}
func test(word string, str string, input string) {
}

// function pour afficher
func ToFind(word string, str string, input string) bool {
	for i := 0; i < len(word); i++ {
		if strings.Contains(string(word[i]), str) {
			fmt.Print(str)
		} else if strings.Contains(string(word[i]), input) {
			return true
		}
	}
	return false
}

func game() {
	input := ""
	args := os.Args[1]
	word := readFile(args)
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	ToFind(word, str, input)
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
