/*
You are given an array a consisting of n (n≥3) positive integers. It is known that in this array, all the numbers except one are the same (for example, in the array [4,11,4,4] all numbers except one are equal to 4

).

Print the index of the element that does not equal others. The numbers in the array are numbered from one.
Input

The first line contains a single integer t
(1≤t≤100). Then t

test cases follow.

The first line of each test case contains a single integer n
(3≤n≤100) — the length of the array a

.

The second line of each test case contains n
integers a1,a2,…,an (1≤ai≤100

).

It is guaranteed that all the numbers except one in the a

array are the same.
Output

For each test case, output a single integer — the index of the element that is not equal to others.
*/

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		var correct int
		if a[0] != a[1] {
			if a[0] != a[2] {
				correct = a[1]
			} else {
				correct = a[0]
			}
		} else {
			correct = a[0]
		}
		for j := 0; j < n; j++ {
			if a[j] != correct {
				fmt.Println(j + 1)
				break
			}
		}
	}
}
