package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	dist := 0

	if len(a) != len(b) {
		return -1, errors.New("strings from different sizes")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist += 1
		}
	}
	return dist, nil
}
