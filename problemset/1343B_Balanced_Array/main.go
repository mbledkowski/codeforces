/*
You are given a positive integer n, it is guaranteed that n is even (i.e. divisible by 2

).

You want to construct the array a
of length n

such that:

    The first n2

elements of a are even (divisible by 2
);
the second n2
elements of a are odd (not divisible by 2
);
all elements of a
are distinct and positive;
the sum of the first half equals to the sum of the second half (∑i=1n2ai=∑i=n2+1nai

    ).

If there are multiple answers, you can print any. It is not guaranteed that the answer exists.

You have to answer t

independent test cases.
Input

The first line of the input contains one integer t
(1≤t≤104) — the number of test cases. Then t

test cases follow.

The only line of the test case contains one integer n
(2≤n≤2⋅105) — the length of the array. It is guaranteed that that n is even (i.e. divisible by 2

).

It is guaranteed that the sum of n
over all test cases does not exceed 2⋅105 (∑n≤2⋅105

).
Output

For each test case, print the answer — "NO" (without quotes), if there is no suitable answer for the given test case or "YES" in the first line and any suitable array a1,a2,…,an
(1≤ai≤109) satisfying conditions from the problem statement on the second line.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	writer := bufio.NewWriter(os.Stdout)
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		if n%4 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			var a []int = make([]int, n)
			for j := 0; j < n/2; j++ {
				a[j] = (j + 1) * 2
			}
			for j := n / 2; j < n; j++ {
				a[j] = (j-n/2)*2 + 1
			}
			a[n-1] += n / 2
			for j := 0; j < n; j++ {
				writer.WriteString(strconv.Itoa(a[j]))
				if j != n-1 {
					writer.WriteString(" ")
				} else {
					writer.WriteString("\n")
				}
			}
		}
		writer.Flush()
	}
}
