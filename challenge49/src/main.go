package main

import (
	"fmt"
	"src/algorithm"
	"strings"
)

func main() {
	fmt.Println("Which approach you want to use bruteForce/wordList?")
	userInput := readUserInput()
	if strings.EqualFold(userInput, "bruteForce") {
		fmt.Println("Enter the Hash to decode, for now it is supported for passwords of length 6")
		hashToDecode := readUserInput()
		if algorithm.BruteForce(hashToDecode) {
			fmt.Println("Thanks for using hash decoder")
		} else {
			fmt.Println("Sorry couldn't decode the hash")
		}
	} else if strings.EqualFold(userInput, "wordList") {
		fmt.Println("Enter the Hash to decode")
		hashToDecode := readUserInput()
		if algorithm.WordList(hashToDecode) {
			fmt.Println("Thanks for using hash decoder")
		} else {
			fmt.Println("Sorry couldn't decode the hash")
		}
	}
}

func readUserInput() string {
	var input string

	fmt.Print("Enter Your Input: ")

	fmt.Scanln(&input)

	fmt.Println("You entered:", input)

	return input
}
