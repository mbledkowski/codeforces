/*
There are _n_ stones on the table in a row, each of them can be red, green or blue. Count the minimum number of stones to take from the table so that any two neighboring stones had different colors. Stones in a row are considered neighboring if there are no other stones between them.

Input

The first line contains integer _n_ (1 ≤ _n_ ≤ 50) — the number of stones on the table.

The next line contains string _s_, which represents the colors of the stones. We'll consider the stones in the row numbered from 1 to _n_ from left to right. Then the _i_\-th character _s_ equals "R", if the _i_\-th stone is red, "G", if it's green and "B", if it's blue.

Output

Print a single integer — the answer to the problem.
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	var result int = 0
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			result++
		}
	}
	fmt.Println(result)
}
