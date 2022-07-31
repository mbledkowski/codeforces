/*
AquaMoon has two binary sequences a and b, which contain only 0 and 1. AquaMoon can perform the following two operations any number of times (a1 is the first element of a, a2 is the second element of a

, and so on):

    Operation 1: if a

contains at least two elements, change a2 to min(a1,a2), and remove the first element of a
.
Operation 2: if a
contains at least two elements, change a2 to max(a1,a2), and remove the first element of a

    .

Note that after a removal of the first element of a
, the former a2 becomes the first element of a, the former a3 becomes the second element of a and so on, and the length of a

reduces by one.

Determine if AquaMoon can make a
equal to b

by using these operations.
Input

The first line contains a single integer t
(1≤t≤2000

) — the number of test cases. Description of test cases follows.

The first line of each test case contains two integers n
, m (1≤n,m≤50, m≤n) — the lengths of a and b

respectively.

The second line of each test case contains a string a
of length n, consisting only 0 and 1

.

The third line of each test case contains a string b
of length m, consisting only 0 and 1

.
Output

For each test case, output "YES" if AquaMoon can change a
to b

by using these options; otherwise, output "NO".

You may print each letter in any case (for example, "YES", "Yes", "yes", "yEs" will all be recognized as a positive answer).
*/

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		var a string
		var b string

		fmt.Scan(&a)
		fmt.Scan(&b)

		for j := 0; j < len(a); j++ {
			if j >= len(b) || a[j] != b[j] {
				if a[0] != b[0] {
					a = a[1:] // remove first element
					j = 0
				} else {
					a = a[0:1] + a[2:] // remove second element
					j = 0
				}
			}
		}
		if a == b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
