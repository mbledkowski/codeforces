/*
Vanya and his friends are walking along the fence of height _h_ and they do not want the guard to notice them. In order to achieve this the height of each of the friends should not exceed _h_. If the height of some person is greater than _h_ he can bend down and then he surely won't be noticed by the guard. The height of the _i_\-th person is equal to _a__i_.

Consider the width of the person walking as usual to be equal to 1, while the width of the bent person is equal to 2. Friends want to talk to each other while walking, so they would like to walk in a single row. What is the minimum width of the road, such that friends can walk in a row and remain unattended by the guard?

Input

The first line of the input contains two integers _n_ and _h_ (1 ≤ _n_ ≤ 1000, 1 ≤ _h_ ≤ 1000) — the number of friends and the height of the fence, respectively.

The second line contains _n_ integers _a__i_ (1 ≤ _a__i_ ≤ 2_h_), the _i_\-th of them is equal to the height of the _i_\-th person.

Output

Print a single integer — the minimum possible valid width of the road.
*/

package main

import "fmt"

func minWidth(n int, h int, a []int) int {
	var min int
	for _, v := range a {
		if v > h {
			min += 2
		} else {
			min++
		}
	}
	return min
}

func main() {
	var n, h int
	fmt.Scan(&n, &h)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(minWidth(n, h, a))
}
