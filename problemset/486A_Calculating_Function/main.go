/*
For a positive integer _n_ let's define a function _f_:

_f_(_n_) =  - 1 + 2 - 3 + .. + ( - 1)_n__n_

Your task is to calculate _f_(_n_) for a given integer _n_.

Input

The single line contains the positive integer _n_ (1 ≤ _n_ ≤ 1015).

Output

Print _f_(_n_) in a single line.
*/

package main

import (
	"fmt"
)

func f(n uint64) int64 {
	if n%2 == 0 {
		return int64(n / 2)
	} else {
		return -int64(n/2 + 1)
	}
}

func main() {
	var n uint64
	fmt.Scan(&n)
	fmt.Println(f(n))
}
