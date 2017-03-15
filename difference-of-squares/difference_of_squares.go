package diffsquares

const testVersion = 1

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func SquareOfSums(n int) int {
	sums := 0
	for i := 1; i <= n; i++ {
		sums += i
	}
	return sums * sums
}

func Difference(n int) int {
	sqrSum := SquareOfSums(n)
	sumSqr := SumOfSquares(n)
	difference := sqrSum - sumSqr
	return difference
}
