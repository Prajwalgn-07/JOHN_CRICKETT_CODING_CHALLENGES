package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Which approach you want to use bruteForce/wordList?")
	userInput := readUserInput()
	if strings.EqualFold(userInput, "bruteForce") {
		fmt.Println("Enter the Hash to decode")
		hashToDecode := readUserInput()
		if bruteForce(hashToDecode) {
			fmt.Println("Thanks for using hash decoder")
		} else {
			fmt.Println("Sorry couldn't decode the hash")
		}
	} else if strings.EqualFold(userInput, "wordList") {
		fmt.Println("Enter the Hash to decode")
		hashToDecode := readUserInput()
		if wordList(hashToDecode) {
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

func bruteForce(testHash string) (isPasswordCracked bool) {
	bruteForceCombinations := getAllBruteForceCombinationsForTheSpecifiedLength(4)
	for _, password := range bruteForceCombinations {
		hash := md5.Sum([]byte(password))
		hashInString := hex.EncodeToString(hash[:])
		if hashInString == testHash {
			fmt.Println("The password for the hash is:", password)
			return true
		}
	}
	return false
}

func wordList(testHash string) (isPasswordCracked bool) {
	wordListCombinations := readWordList()
	for _, password := range wordListCombinations {
		hash := md5.Sum([]byte(password))
		hashInString := hex.EncodeToString(hash[:])
		if hashInString == testHash {
			fmt.Println("The password for the hash is:", password)
			return true
		}
	}
	return false
}

func readWordList() (combinations []string) {
	// Open the text file
	file, err := os.Open("word-list.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the lines
	var lines []string

	// Read lines and append to the slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	return lines
}
