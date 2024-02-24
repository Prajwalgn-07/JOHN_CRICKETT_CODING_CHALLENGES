package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// func preComputHashesFromWordList(wordListFile string) {
// 	listOfWords := ReadWordList(wordListFile)
// 	for _, password := range listOfWords {
// 		hashInString := GetHash(password)
// 		WriteHashPasswordToRainBowTableJson(hashInString, password, "word-list")
// 	}
// }

func ReadWordList(fileName string) (combinations []string) {
	// Open the text file
	file, err := os.Open(fileName)
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

func WriteHashPasswordToRainBowTableJson(hash, password string, source string) error {
	// Hardcoded file name for now
	fileName := "rainbowTable.json"

	// Read existing JSON data from file
	byteValue, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into a map
	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return err
	}

	// Get jsonData data
	jsonDataOfSource, ok := data[source].(map[string]interface{})
	if !ok {
		jsonDataOfSource = make(map[string]interface{})
	}

	// Add the new hash-password pair to jsonData
	jsonDataOfSource[hash] = password

	// Update the data map with the modified bruteForceTable
	data[source] = jsonDataOfSource

	// Marshal the updated data back to JSON
	updatedData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return err
	}

	return nil
}

// func preComputHashesFromBruteForceMethod(lengthOfThehashes int) {
// 	bruteForceCombinations := getAllBruteForceCombinationsForTheSpecifiedLength(lengthOfThehashes)
// 	for _, password := range bruteForceCombinations {
// 		hashInString := GetHash(password)
// 		WriteHashPasswordToRainBowTableJson(hashInString, password, "bruteForceTable")
// 	}
// }
