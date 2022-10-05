package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func readFile() {
	data, _ := ioutil.ReadFile("words/words.txt")
	file := string(data)
	temp := strings.Split(file, "\n")
	l := len(temp)
	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(l)
	fmt.Print(j)
}

func main() {
	readFile
}
