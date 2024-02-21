package main

func getAllBruteForceCombinationsForTheSpecifiedLength(length int) []string {
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
