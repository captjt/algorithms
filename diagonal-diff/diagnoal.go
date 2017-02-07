package diagonalDiff

import (
	"fmt"
	"math"
)

func read() (int, [][]int) {
	var n int
	fmt.Scan(&n)
	s := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := range row {
			fmt.Scan(&row[j])
		}
		s = append(s, row)
	}
	return n, s
}

func main() {
	// variables to track diagonal totals
	var primTotal, secTotal int

	size, numbers := read()

	primary := 0
	secondary := size - 1

	for i := 0; i < size; i++ {
		primTotal += numbers[i][primary]
		secTotal += numbers[i][secondary]

		primary++
		secondary--
	}

	pf := float64(primTotal)
	sf := float64(secTotal)

	fmt.Println(math.Abs(pf - sf))
}
