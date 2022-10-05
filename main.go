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
	line := 0
	temp := strings.Split(file, "\n")
	l := len(temp)
	for _,item := range temp {
		fmt.Println(line,item)
		line++
	}
	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(l)
	fmt.Print(j)
}

func main() {
	readFile()
}
