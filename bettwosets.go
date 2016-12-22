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
	_ = sizes
	a := convertToIntArray(second)
	b := convertToIntArray(third)
	comb := append(a, b...)

	validator := checkDenominator(a, b)

	factorMap := map[int]int{}

	if validator == true {
		for _, i := range comb {
			factorArray := numberFactors(i)
			for _, j := range factorArray {
				factorMap[j] += 1
			}
		}
	} else {
		fmt.Println(0)
		return
	}

	count := countFactors(factorMap, a)
	fmt.Println(count)
}

func convertToIntArray(s string) []int {
	array := make([]int, 0)
	trim := strings.TrimSpace(s) // hackerrank adds newline spacing must trim
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

func checkDenominator(a []int, b []int) bool {
	for i := range b {
		for j := range a {
			if b[i]%a[j] != 0 {
				return false
			}
		}
	}
	return true
}

func numberFactors(n int) []int {
	array := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			array = append(array, i)
		}
	}
	return array
}

func countFactors(m map[int]int, a []int) int {
	array := make([]int, 0)
	for key, val := range m {
		for _, i := range a {
			if val > 1 && key >= i {
				array = append(array, val)
			}
		}
	}
	return len(array)
}
