package triangle

import "math"

const testVersion = 3
const (
	NaT Kind = 0
	Equ Kind = 1
	Iso Kind = 2
	Sca Kind = 3
)

type Kind int

func KindFromSides(a, b, c float64) Kind {
	if existSide(a) && existSide(b) && existSide(c) && existTriangle(a, b, c) {
		if a == b && b == c {
			return Equ
		} else if a != b && b != c && c != a {
			return Sca
		}
		return Iso
	}
	return NaT
}

func existSide(l float64) bool {
	if l <= 0 || math.IsNaN(l) || math.IsInf(l, 0) {
		return false
	}
	return true
}

func existTriangle(a, b, c float64) bool {
	if a > b+c || b > a+c || c > a+b {
		return false
	}
	return true
}
