package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func ChooseWord() {
	data, _ := ioutil.ReadFile("words/words.txt")
	file := string(data)
	temp := strings.Split(file, "\n")
	l := len(temp)
}

func main() {
	data, _ := ioutil.ReadFile("words/words.txt")
	file := string(data)
	temp := strings.Split(file, "\n")
	l := len(temp)
	ChooseWord()
	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(l)
	fmt.Print(j)
}
