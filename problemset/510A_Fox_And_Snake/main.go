/*
Fox Ciel starts to learn programming. The first task is drawing a fox! However, that turns out to be too hard for a beginner, so she decides to draw a snake instead.

A snake is a pattern on a _n_ by _m_ table. Denote _c_\-th cell of _r_\-th row as (_r_, _c_). The tail of the snake is located at (1, 1), then it's body extends to (1, _m_), then goes down 2 rows to (3, _m_), then goes left to (3, 1) and so on.

Your task is to draw this snake for Fox Ciel: the empty cells should be represented as dot characters ('.') and the snake cells should be filled with number signs ('#').

Consider sample tests in order to understand the snake pattern.

Input

The only line contains two integers: _n_ and _m_ (3 ≤ _n_, _m_ ≤ 50).

_n_ is an odd number.

Output

Output _n_ lines. Each line should contain a string consisting of _m_ characters. Do not output spaces.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println(strings.Repeat("#", m))
		} else {
			if i%4 == 1 {
				fmt.Println(strings.Repeat(".", m-1) + "#")
			} else {
				fmt.Println("#" + strings.Repeat(".", m-1))
			}
		}
	}
}
