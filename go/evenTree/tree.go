// challenge :: https://www.hackerrank.com/challenges/even-tree

package main

import "fmt"

func main() {
	var a, b, x, y int
	tree := make(map[int][]int)

	fmt.Scanf("%v %v\n", &a, &b)

	for i := 0; i < a; i++ {
		fmt.Scanf("%v %v\n", &x, &y)

		if _, ok := tree[y]; ok {
			tree[y] = append(tree[y], x)
		} else {
			tree[y] = append(tree[y], x)
		}

		fmt.Println(tree)
	}
}
