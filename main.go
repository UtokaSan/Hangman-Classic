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

// faire une boucle pour répéter l'input à chaque fois
func ToFind(word string, str string, input string) {
	for i := 0; i < len(word); i++ {
		strToInt := string(word[i])
		strToInt = "_"
		fmt.Print(strToInt)
		fmt.Scanf("%s", input)
		if string(word[i]) == str {
			fmt.Print(str)
		}
		if string(word[i]) == input {
			fmt.Print(input)
		}
	}
}

func game() {
	input := ""
	args := os.Args[1]
	word := readFile(args)
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
	fmt.Scanf("%s", &input)
	ToFind(word, str, input)
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
