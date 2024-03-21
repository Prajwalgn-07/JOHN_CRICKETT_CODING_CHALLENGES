package algorithm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"src/utils"
)

func BruteForce(hash string) (isPasswordCracked bool) {
	bruteForceCombinations := GetAllParallelCombinationsForTheSpecifiedLength(6)
	for _, password := range bruteForceCombinations {
		hashInString := GetHash(password)
		if hashInString == hash {
			fmt.Println("The password for the hash is:", password)
			return true
		}
	}
	return false
}

func WordList(hash string) (isPasswordCracked bool) {
	wordListCombinations := utils.ReadWordList("word-list.txt")
	for _, password := range wordListCombinations {
		hashInString := GetHash(password)
		if hashInString == hash {
			fmt.Println("The password for the hash is:", password)
			return true
		}
	}
	return false
}

func ReadWordList(s string) {
	panic("unimplemented")
}

func GetHash(word string) string {
	hash := md5.Sum([]byte(word))
	hashInString := hex.EncodeToString(hash[:])
	return hashInString
}
