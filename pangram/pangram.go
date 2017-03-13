package pangram

import "strings"

const testVersion = 1

func IsPangram(str string) bool {
	alphabet := strings.Split("abcdefghijklmnopqrstuvxwyz", "")
	inpuText := strings.ToLower(str)

	for _, letter := range alphabet {
		if strings.ContainsAny(letter, inpuText) {
			continue
		} else {
			return false
		}
	}
	return true
}
