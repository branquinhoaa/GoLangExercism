package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

var pattern = regexp.MustCompile(`(\b\w)|[^A-Z\s][A-Z]`)

func Abbreviate(str string) string {
	array := pattern.FindAllString(str, -1)

	for i, char := range array {
		array[i] = strings.ToUpper(char[len(char)-1:])
	}
	return strings.Join(array, "")
}
