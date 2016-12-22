// Link : https://www.hackerrank.com/challenges/between-two-sets
// Requirements
// All elements in A are factors of x.
// x is a factor of all elements in B.

// INPUT
// 2 3
// 2 4
// 16 32 96
// OUTPUT
// 3
// EXPLANATION
// The integers that are between A = {2,4} and B = {16, 32, 96} are 4, 8, and 16.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first, _ := reader.ReadString('\n')
	second, _ := reader.ReadString('\n')
	third, _ := reader.ReadString('\n')

	sizes := convertToIntArray(first)
	a := convertToIntArray(second)
	b := convertToIntArray(third)

	fmt.Println(sizes)
	fmt.Println(a)
	fmt.Println(b)
}

func convertToIntArray(s string) []int {
	array := make([]int, 0)
	trim := strings.TrimSpace(s)
	str := strings.Split(trim, " ")

	for _, i := range str {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		array = append(array, j)
	}
	return array
}
