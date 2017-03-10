package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(n int) string {
	sentence := ""
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		return strconv.Itoa(n)
	}
	if n%3 == 0 {
		sentence += "Pling"
	}
	if n%5 == 0 {
		sentence += "Plang"
	}
	if n%7 == 0 {
		sentence += "Plong"
	}
	return sentence
}
