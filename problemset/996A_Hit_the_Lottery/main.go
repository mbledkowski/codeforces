/*
Allen has a LOT of money. He has n dollars in the bank. For security reasons, he wants to withdraw it in cash (we will not disclose the reasons here). The denominations for dollar bills are 1, 5, 10, 20, 100

. What is the minimum number of bills Allen could receive after withdrawing his entire balance?
Input

The first and only line of input contains a single integer n
(1≤n≤109

).
Output

Output the minimum number of bills that Allen could receive.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var count int
	count += int(math.Floor(float64(n) / 100))
	n = n % 100
	for n > 0 {
		if n >= 20 {
			count += int(math.Floor(float64(n) / 20))
			n = n % 20
		} else if n >= 10 {
			count += int(math.Floor(float64(n) / 10))
			n = n % 10
		} else if n >= 5 {
			count += int(math.Floor(float64(n) / 5))
			n = n % 5
		} else if n >= 1 {
			count += int(math.Floor(float64(n) / 1))
			n = n % 1
		}
	}
	fmt.Println(count)
}
