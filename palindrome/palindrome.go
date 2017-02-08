// premise of this is to validate strings to be a palindrome
// this means a string is spelled the same way backwards as it is regularly

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"blutulb", "test", "  ", "racecar", "do  od", "lefties"}

	for _, v := range s {
		check := validatePalidrome(v)
		fmt.Printf("Word: >>> '%s' <<<\nResult: >>> %t <<<\n\n", v, check)
	}
}

func validatePalidrome(s string) (b bool) {
	word := strings.Split(s, "")
	first := 0
	last := len(word) - 1

	for {
		if last <= first {
			break
		}

		if word[first] == word[last] {
			b = true
			first++
			last--
		} else {
			b = false
			break
		}
	}

	return
}
