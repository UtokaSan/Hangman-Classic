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
func game() {
	input := ""
	word := readFile("words/words.txt")
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	fmt.Scanf("%s", &input)
	for _, el := range word { // faire une boucle tant que ce n'est pas le mot demandé
		strToInt := string(el)
		if strToInt == str {
			fmt.Print(strToInt)
		} else if strToInt == input {
			fmt.Print(strToInt)
		} else {
			el = 95
			strToInt2 := string(el)
			fmt.Print(strToInt2)
		}
	}
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
