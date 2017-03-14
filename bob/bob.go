package bob // package name must match the package name in bob_test.go
import (
	"strings"
)

const testVersion = 2 // same as targetTestVersion

func Hey(str string) string {
	words := strings.TrimSpace(str)
	if strings.ToUpper(words) == words && strings.ToLower(words) != words {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(words, "?") {
		return "Sure."
	} else if words == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
