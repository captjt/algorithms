package main

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
	var a, res int
	fmt.Scanf("%v\n", &a)

	numbers, err := read(a)
	if err != nil {
		panic(err)
	}

	for _, val := range numbers {
		res += val
	}

	fmt.Println(res)
}
