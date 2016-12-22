package main

// Link : https://www.hackerrank.com/challenges/mini-max-sum

// OBJECTIVE:
// Given five positive integers, find the minimum and maximum values that can be
//  calculated by summing exactly four of the five integers. Then print the respective
//  minimum and maximum values as a single line of two space-separated long integers.

// INPUT
// A single line of five space-separated integers.
// OUTPUT
// Print two space-separated long integers denoting the respective minimum and maximum
//  values that can be calculated by summing exactly four of the five integers. (The output
//  can be greater than 32 bit integer.)

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	array := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	s := strings.Split(text, " ")

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		array = append(array, j)
	}

	sort.Ints(array)

	length := len(array)

	min := addSilce(array[0:4])
	max := addSilce(array[length-4 : length])

	fmt.Printf("%d %d", min, max)
}

func addSilce(array []int) int {
	total := 0

	for _, i := range array {
		total += i
	}

	return total
}
