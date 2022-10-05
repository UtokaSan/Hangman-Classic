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
	line := 0
	temp := strings.Split(file, "\n")
	l := len(temp)
	for _, item := range temp {
		fmt.Println("[", line, "]\t", item)
		line++
	}
	fmt.Println(l)
}

func main() {
	data, _ := ioutil.ReadFile("words/words.txt")
	file := string(data)
	temp := strings.Split(file, "\n")
	l := len(temp)
	ChooseWord()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(l))
}
