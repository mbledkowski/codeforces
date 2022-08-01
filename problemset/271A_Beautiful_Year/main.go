/*
It seems like the year of 2013 came only yesterday. Do you know a curious fact? The year of 2013 is the first year after the old 1987 with only distinct digits.

Now you are suggested to solve the following problem: given a year number, find the minimum year number which is strictly larger than the given one and has only distinct digits.
Input

The single line contains integer y (1000 ≤ y ≤ 9000) — the year number.
Output

Print a single integer — the minimum year number that is strictly larger than y and all it's digits are distinct. It is guaranteed that the answer exists.
*/

package main

import "fmt"

func isDistinct(y int) bool {
	var digits [10]bool
	for y > 0 {
		if digits[y%10] {
			return false
		}
		digits[y%10] = true
		y /= 10
	}
	return true
}

func main() {
	var y int
	fmt.Scan(&y)
	for {
		y++
		if isDistinct(y) {
			break
		}
	}
	fmt.Println(y)
}
