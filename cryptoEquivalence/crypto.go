// Crypto-Equivalence
// Two strings are crypto-equivalent if there exists some reversible mapping of letters to letters
//   that transform the first string into the second string.
// For example, 'cat' and 'dog' are crypto-equivalent since the mapping (c -> d, a -> o, t -> g)
//   transforms 'cat' into 'dog'
// However, 'cat' and 'all' are not crypto-equivalent since (c -> a, a -> l, t -> l) is not a
//   reversible operation, and there is now way to transform 'cat' into 'all'.
// Write a function that groups a list of strings into a list of lists, where each list contains all
//   strings that are crypto-equivalent to each other.
// Example grouping: ['dog', 'cat', 'all', 'dad', 'xyz'] => [['dog', 'cat', 'xyz'], ['all'], ['dad']]

package main

import (
	"sort"
	"strings"
)

type groups struct {
	mapping []int
	strings []string
}

func main() {
	var g []groups

	s := []string{"dog", "cat", "all", "dad", "xyz"}

	for _, v := range s {
		c := makeCrypt(v)

		n := convertMapping(c)

		reconcileCrypto(&g, n, v)
	}
}

func makeCrypt(s string) map[string]int {
	res := map[string]int{}

	f := func(key string) bool { _, ok := res[key]; return ok }

	a := strings.Split(s, "")

	for _, v := range a {
		switch {
		case f(v):
			res[v]++
		default:
			res[v] = 1
		}
	}
	return res
}

func convertMapping(c map[string]int) (res []int) {
	for _, v := range c {
		res = append(res, v)
	}

	sort.Ints(res)

	return
}

func reconcileCrypto(g *[]groups, n []int, s string) {
	for key := range g {
        res := compareArrays(g[key].mappings, n)
        if
	}
}

func compareArrays(a, b []int) (c bool) {
	if len(a) != len(b) {
		return
	}

	for i, v := range a {
		if v != b[i] {
			c = false
		}
		c = true
	}
	return
}
