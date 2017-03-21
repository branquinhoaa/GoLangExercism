package queenattack

import "errors"

const testVersion = 2

func CanQueenAttack(queen1 string, queen2 string) (bool, error) {
	if ErrorCases(queen1, queen2) {
		return false, errors.New("error")
	}
	if RightPosition(queen1, queen2) {
		return true, nil
	}

	return false, nil
}

func ErrorCases(a string, b string) bool {
	if len(a) != 2 || len(b) != 2 || a == b {
		return true
	} else if a[1] > 56 || a[1] < 49 || a[0] > 104 {
		return true
	} else if b[1] > 56 || b[1] < 49 || b[0] > 104 {
		return true
	}
	return false
}

func RightPosition(a string, b string) bool {
	for i := 0; i < 2; i++ {
		if a[i] == b[i] && a[1-i] != b[1-i] {
			return true
		}
	}
	if a[0]-b[0] == a[1]-b[1] || a[0]-b[0] == b[1]-a[1] {
		return true
	}
	return false
}
