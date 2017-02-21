// challenge :: https://www.hackerrank.com/challenges/staircase

package staircase

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	fmt.Scanf("%v\n", &a)

	for i := 0; i < a; i++ {
		hashes := i + 1
		spaces := a - hashes
		fmt.Println(strings.Repeat(" ", spaces) + strings.Repeat("#", hashes))
	}
}
