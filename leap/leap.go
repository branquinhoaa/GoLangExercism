package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}
