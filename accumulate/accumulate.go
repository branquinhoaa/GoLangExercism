package accumulate

const testVersion = 1

func Accumulate(t []string, f func(string) string) []string {
	for i, element := range t {
		t[i] = f(element)
	}
	return t
}
