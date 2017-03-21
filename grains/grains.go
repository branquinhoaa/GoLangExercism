package grains

import (
	"errors"
	"math"
)

const testVersion = 1

func Square(square int) (uint64, error) {
	if square > 64 || square < 1 {
		return 0, errors.New("error")
	}
	result := math.Pow(2, float64(square-1))
	return uint64(result), nil
}

func Total() (total uint64) {
	for i := 1; i < 65; i++ {
		sum, _ := Square(i)
		total += sum
	}
	return total
}
