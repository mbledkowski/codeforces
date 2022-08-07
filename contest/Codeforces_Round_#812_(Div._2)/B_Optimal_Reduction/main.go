/*
Consider an array a of n

positive integers.

You may perform the following operation:

    select two indices l

and r (1≤l≤r≤n
), then
decrease all elements al,al+1,…,ar
by 1

    .

Let's call f(a)
the minimum number of operations needed to change array a into an array of n

zeros.

Determine if for all permutations†
b of a, f(a)≤f(b)

is true.

†
An array b is a permutation of an array a if b consists of the elements of a in arbitrary order. For example, [4,2,3,4] is a permutation of [3,2,4,4] while [1,2,2] is not a permutation of [1,2,3]

.
Input

The first line contains a single integer t
(1≤t≤104

) — the number of test cases.

The first line of each test case contains a single integer n
(1≤n≤105) — the length of the array a

.

The second line contains n
integers a1,a2,…,an (1≤ai≤109) — description of the array a

.

It is guaranteed that the sum of n
over all test cases does not exceed 105

.
Output

For each test case, print "YES" (without quotes) if for all permutations b
of a, f(a)≤f(b)

is true, and "NO" (without quotes) otherwise.

You can output "YES" and "NO" in any case (for example, strings "yEs", "yes" and "Yes" will be recognized as a positive response).
*/

package main

import (
	"bufio"
	"os"
	// "sort"
	"strconv"
	"strings"
)

// func f(a []int, n int, iteration int, maxIteration int) int {
// 	if iteration > maxIteration && maxIteration != -1 {
// 		return -1
// 	}
// 	var l, r int = -1, -1
// 	for i := 0; i < n; i++ {
// 		if a[i] != 0 {
// 			l = i
// 			r = n - 1
// 			break
// 		}
// 	}
// 	if l == -1 {
// 		return iteration
// 	}
// 	for i := l + 1; i < n; i++ {
// 		if a[i] == 0 {
// 			r = i - 1
// 			break
// 		}
// 	}
// 	for i := l; i <= r; i++ {
// 		a[i]--
// 	}
// 	return f(a, n, iteration+1, maxIteration)
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		astr := strings.Split(scanner.Text(), " ")
		a := make([]int, n)
		// asorted := make([]int, n)
		for j := 0; j < n; j++ {
			a[j], _ = strconv.Atoi(astr[j])
			// asorted[j] = a[j]
		}
		// sort.Ints(asorted)
		// fasorted := f(asorted, n, 0, -1)
		// fa := f(a, n, 0, fasorted)
		// if fa == -1 {
		// 	writer.WriteString("NO\n")
		// } else {
		// 	writer.WriteString("YES\n")
		// }
		var min = a[0]
		var mini = 0
		var res bool = true
		var res2 bool = true
		for j := 1; j < n; j++ {
			if a[j] > min && min != a[0] {
				res = false
				res2 = false
				break
			}
			if a[j] < min {
				min = a[j]
				mini = j
			}
		}
		for j := 0; j < mini; j++ {
			if a[j] > min {
				res = false
				break
			}
		}
		for j := mini; j < n; j++ {
			if a[j] > min {
				res2 = false
				break
			}
		}

		if res || res2 {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
	writer.Flush()
}
