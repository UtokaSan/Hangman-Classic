package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func Random(nbr int) int {
	//rand.Seed(time.Now().UnixNano()) Pour random tout le temps
	return rand.Intn(nbr)
}
func readFile() {
	data, _ := ioutil.ReadFile("words/words.txt")
	str := string(data)
	split := strings.Split(str, "\n")
	random := Random(len(split))
	word := split[random]
	fmt.Print(word)
}

func main() {
	readFile()
}
