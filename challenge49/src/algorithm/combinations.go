package algorithm

func GetAllBruteForceCombinationsForTheSpecifiedLength(length int) []string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789@#!()"
	var combinations []string
	for i := 1; i <= length; i++ {
		generateCombinations(alphabet, length, "", &combinations)
	}
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
