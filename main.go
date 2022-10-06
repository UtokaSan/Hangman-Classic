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
	word := readFile("words/words.txt")
	len := len(word)/2 - 1
	str := string(word[len])
	fmt.Print(str)
}
func readFile(fileAccess string) string {
	data, _ := ioutil.ReadFile(fileAccess)
	str := string(data)
	split := strings.Split(str, "\n")
	random := Random(len(split))
	word := split[random]
	fmt.Println(word)
	return word
}

func main() {
	test()
}
