package etl

import "strings"

const testVersion = 1

func Transform(etl map[int][]string) map[string]int {
	result := make(map[string]int)
	for key, value := range etl {
		for i := 0; i < len(value); i++ {
			newKey := strings.ToLower(value[i])
			result[newKey] = key
		}
	}
	return result
}
