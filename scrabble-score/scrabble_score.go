package scrabble

import "strings"

const testVersion = 4

func Score(s string) int {
	word := strings.ToUpper(s)
	sum := 0
	alphabet := map[int]string{1: "AEIOULNRST", 2: "DG", 3: "BCMP", 4: "FHVWY", 5: "K", 8: "JX", 10: "QZ"}
	for i := 0; i < len(word); i++ {
		letter := word[i]
		for key, value := range alphabet {
			if strings.ContainsAny(value, string(letter)) {
				sum += key
			}
		}
	}
	return sum
}
