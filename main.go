package main

import (
	"fmt"
	"io"
	"os"
)

var path = "C:/Users/jorda/OneDrive/Documents/GitHub/Hangman/words/words.txt"

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				panic(err)
			}
		}
		if n == 0 {
			break
		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("le mot sera :", string(text))
}

func main() {
	readFile()
}
