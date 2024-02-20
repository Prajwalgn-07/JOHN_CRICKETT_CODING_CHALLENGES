package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	testHash := "7a95bf926a0333f57705aeac07a362a2"
	combinations := getAllCombinationsForTheSpecifiedLength(4)

	for _, password := range combinations {
		hash := md5.Sum([]byte(password))
		hashInString := hex.EncodeToString(hash[:])
		if hashInString == testHash {
			fmt.Println("The password for the hash is:", password)
			break
		}
	}

	fmt.Println("Search complete.")
}

func getAllCombinationsForTheSpecifiedLength(length int) []string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var combinations []string
	generateCombinations(alphabet, length, "", &combinations)
	return combinations
}

func generateCombinations(alphabet string, length int, prefix string, combinations *[]string) {
	if length == 0 {
		*combinations = append(*combinations, prefix)
		return
	}
	for _, char := range alphabet {
		generateCombinations(alphabet, length-1, prefix+string(char), combinations)
	}
}
