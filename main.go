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
func test() {
	input := 0
	word := readFile("words/words.txt")
	fmt.Println(word)
	str := string(word[len(word)/2-1])
	fmt.Println(str)
	for _, el := range word {
		strToInt := string(el)
		if strToInt == str {
			fmt.Print(strToInt)
		} else {
			el = 95
			strToInt2 := string(el)
			fmt.Print(strToInt2)
		}
	}
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
