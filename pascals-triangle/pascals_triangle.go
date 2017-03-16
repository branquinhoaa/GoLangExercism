package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	if n == 1 {
		one := [][]int{}
		one = append(one, []int{1})
		return one
	}
	return CreateRow(n)
}

func CreateRow(n int) [][]int {
	all := [][]int{}
	for j := 2; j <= n; j++ {
		if j == 2 {
			one := []int{1}
			two := []int{1, 1}
			all = append(all, one)
			all = append(all, two)
		} else {
			newRow := []int{1}
			for i := 1; i < j-1; i++ {
				digit := all[j-2][i-1] + all[j-2][i]
				newRow = append(newRow, digit)
			}
			newRow = append(newRow, 1)
			all = append(all, newRow)
		}
	}
	return all
}
