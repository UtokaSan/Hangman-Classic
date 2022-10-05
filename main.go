package main

import "fmt"

func Menu() {
	var answer int

	fmt.Println("----- Menu -----")
	fmt.Println("1. Easy")
	fmt.Println("2. Hard\n")
	fmt.Println("Choose a number difficulty")
	fmt.Scanf("%d", &answer)
	chooseDifficulty(answer)
}

func chooseDifficulty(answer int) {
	switch answer {
	case 1:
		Easy()
	case 2:
		Hard()
	default:
		fmt.Print("Not valid number")
	}
}

func Easy() {
	fmt.Print("Easy")
}

func Hard() {
	fmt.Print("Hard")
}

func main() {
	Menu()
}
