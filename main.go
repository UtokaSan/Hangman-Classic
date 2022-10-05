package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func Random(nbr int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(nbr)
}
func readFile() {
	data, _ := ioutil.ReadFile("words/words.txt")
	str := string(data)
	split := strings.Split(str, "\n")
	random := Random(len(split))
	fmt.Print(split[random])
}

func main() {
	readFile()
}
