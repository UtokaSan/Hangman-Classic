package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type HangmanData struct { // Faire quelque chose avec la structure normalement c'est pour aider
	word             string
	toFind           string
	attempts         int
	hangmanPositions [10]string
}

func Random(nbr int) int {
	return rand.Intn(nbr)
}
func test() {
	input := 0
	word := readFile("words/words.txt")
	fmt.Println(word)
	str := string(word[len(word)/2-1]) // transforme en string et prends la longueur du mot puis fais un calculs
	fmt.Println(str)
	fmt.Scanf("%s", &input)
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
	test()
}
