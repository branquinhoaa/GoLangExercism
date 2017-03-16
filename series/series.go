package slice

const testVersion = 1

func All(n int, s string) []string {
	var all []string
	for i := 0; i <= len(s)-n; i++ {
		stop := i + n
		substr := s[i:stop]
		all = append(all, substr)
	}
	return all
}

func UnsafeFirst(n int, s string) string {
	start := 0
	stop := start + n
	substring := s[start:stop]
	return substring
}
