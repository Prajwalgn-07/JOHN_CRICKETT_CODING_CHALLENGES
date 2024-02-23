package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Which approach you want to use bruteForce/wordList?")
	userInput := readUserInput()
	if strings.EqualFold(userInput, "bruteForce") {
		fmt.Println("Enter the Hash to decode")
		hashToDecode := readUserInput()
		if BruteForce(hashToDecode) {
			fmt.Println("Thanks for using hash decoder")
		} else {
			fmt.Println("Sorry couldn't decode the hash")
		}
	} else if strings.EqualFold(userInput, "wordList") {
		fmt.Println("Enter the Hash to decode")
		hashToDecode := readUserInput()
		if WordList(hashToDecode) {
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
