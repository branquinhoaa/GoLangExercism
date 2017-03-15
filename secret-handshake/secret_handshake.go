package secret

import "sort"

const testVersion = 1

func Handshake(n uint) []string {
	var handSk []string
	revert := true
	greeting := []string{"wink", "double blink", "close your eyes", "jump"}
	for n > 0 {
		switch {
		case n >= 16:
			revert = false
			n -= 16
		case n >= 8:
			handSk = append(handSk, greeting[3])
			n -= 8
		case n >= 4:
			handSk = append(handSk, greeting[2])
			n -= 4
		case n >= 2:
			handSk = append(handSk, greeting[1])
			n -= 2
		case n == 1:
			handSk = append(handSk, greeting[0])
			n -= 1
		}
	}
	if revert == true {
		return revertArray(handSk)
	}
	return handSk
}

func revertArray(arr []string) []string {
	var strSlice sort.StringSlice = arr
	sort.Sort(sort.Reverse(strSlice))
	return strSlice
}
