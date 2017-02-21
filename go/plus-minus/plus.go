// challenge :: https://www.hackerrank.com/challenges/plus-minus

package plus

import "fmt"

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func main() {
	var a int
	var posRes, negRes, zeroRes float64
	fmt.Scanf("%v\n", &a)

	numbers, err := read(a)
	if err != nil {
		panic(err)
	}

	for _, v := range numbers {
		if v > 0 {
			posRes++
		} else if v < 0 {
			negRes++
		} else if v == 0 {
			zeroRes++
		}
	}

	fa := float64(a)

	fmt.Printf("%.5f\n", posRes/fa)
	fmt.Printf("%.5f\n", negRes/fa)
	fmt.Printf("%.5f\n", zeroRes/fa)
}
