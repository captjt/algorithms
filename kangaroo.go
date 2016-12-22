package main

// Link : https://www.hackerrank.com/challenges/kangaroo

// PURPOSE:
// There are two kangaroos on an x-axis ready to jump in the positive direction
//  (i.e, toward positive infinity). The first kangaroo starts at location x1 and
//  moves at a rate of x2 meters per jump. The second kangaroo starts at location v1
//  and moves at a rate of v2 meters per jump. Given the starting locations and
//  movement rates for each kangaroo, can you determine if they'll ever land at
//  the same location at the same time?

// INPUT:
// A single line of four space-separated integers denoting the respective values
//  of x1, x2, v1, and v2.
// SAMPLE INPUT
// a. 0 3 4 2
// b. 0 2 5 3
// SAMPLE OUTPUT
// a. YES
// b. NO

import "fmt"

type pair struct {
	x, y int
}

type double struct {
	u, w pair
}

func main() {
	var a, b, c, d int
	if _, err := fmt.Scan(&a, &b, &c, &d); err != nil {
		fmt.Println("  Scan for a, b, c & d failed, due to ", err)
	}

	doub := new(double)
	doub.u.x, doub.u.y = a, b
	doub.w.x, doub.w.y = c, d

	response := kangaroo(doub)
	fmt.Println(response)
}

func kangaroo(d *double) string {
	if d.w.y >= d.u.y && d.w.x >= d.u.x {
		return "NO"
	} else if d.u.x == d.w.x {
		return "YES"
	} else if d.u.x >= d.w.x && d.u.y >= d.w.y {
		return "NO"
	} else {
		d.u.x += d.u.y
		d.w.x += d.w.y
		return kangaroo(d)
	}
}
