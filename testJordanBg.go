package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("words/hangman.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 1)
	fmt.Print(buf)

	for {
		n, err := reader.Read(buf)
		if err != nil {

			if err != io.EOF {

				log.Fatal(err)
			}

			break
		}

		fmt.Print(string(buf[0:n]))
	}

	fmt.Println()
}
