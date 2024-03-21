package algorithm

import (
	"math"
	"sync"
)

func GetAllParallelCombinationsForTheSpecifiedLength(length int) []string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	combinations := make([]string, 0, int(math.Pow(float64(len(alphabet)), float64(length))))

	var wg sync.WaitGroup
	var mu sync.Mutex

	generateParallelCombinations(&wg, &mu, alphabet, length, "", &combinations)

	wg.Wait()

	return combinations
}

func generateParallelCombinations(wg *sync.WaitGroup, mu *sync.Mutex, alphabet string, length int, prefix string, combinations *[]string) {
	if length == 0 {
		mu.Lock()
		*combinations = append(*combinations, prefix)
		mu.Unlock()
		return
	}

	for _, char := range alphabet {
		wg.Add(1)
		go func(char rune) {
			defer wg.Done()
			generateParallelCombinations(wg, mu, alphabet, length-1, prefix+string(char), combinations)
		}(char)
	}
}
