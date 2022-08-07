/*
You are given three integers x,y and n. Your task is to find the maximum integer k such that 0≤k≤n that kmodx=y, where mod

is modulo operation. Many programming languages use percent operator % to implement it.

In other words, with given x,y
and n you need to find the maximum possible integer from 0 to n that has the remainder y modulo x

.

You have to answer t
independent test cases. It is guaranteed that such k

exists for each test case.
Input

The first line of the input contains one integer t
(1≤t≤5⋅104) — the number of test cases. The next t

lines contain test cases.

The only line of the test case contains three integers x,y
and n (2≤x≤109; 0≤y<x; y≤n≤109

).

It can be shown that such k

always exists under the given constraints.
Output

For each test case, print the answer — maximum non-negative integer k
such that 0≤k≤n and kmodx=y. It is guaranteed that the answer always exists.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		var k int
		k = n
		if y > n%x {
			k = n - n%x - x + y
		} else {
			k = n - n%x + y
		}
		writer.WriteString(strconv.Itoa(k) + "\n")
	}
	writer.Flush()
}
