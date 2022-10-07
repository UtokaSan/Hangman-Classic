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
func game() {
	input := ""
	args := os.Args[1]
	word := readFile(args)
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println("Good luck, you have 10 attempts")
	fmt.Print("\nChoose : ")
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

func Atoi(s string) int {
	rune := []rune(s)
	stringToInt := 0
	for i := 0; i < len(rune); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stringToInt = stringToInt * 10
			stringToInt = stringToInt + int(rune[i]) - 48
		}
		if s[i] == ' ' {
			return 0
		}
	}
	if s[0] == '-' {
		stringToInt = stringToInt * -1
	}
	if s[0] == '+' && s[1] == '+' || s[0] == '-' && s[1] == '-' {
		return 0
	}
	return stringToInt
}

func main() {
	game()
}
