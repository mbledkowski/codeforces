/*
One way to create a task is to learn from math. You can generate some random math statement or modify some theorems to get something new and build a new task from that.

For example, there is a statement called the "Goldbach's conjecture". It says: "each even number no less than four can be expressed as the sum of two primes". Let's modify it. How about a statement like that: "each integer no less than 12 can be expressed as the sum of two composite numbers." Not like the Goldbach's conjecture, I can prove this theorem.

You are given an integer _n_ no less than 12, express it as a sum of two composite numbers.

Input

The only line contains an integer _n_ (12 ≤ _n_ ≤ 106).

Output

Output two composite integers _x_ and _y_ (1 < _x_, _y_ < _n_) such that _x_ + _y_ = _n_. If there are multiple solutions, you can output any of them.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fmt.Scan(&n)
	for i := n / 2; i < n-1; i++ {
		if !big.NewInt(i).ProbablyPrime(0) {
			for j := n / 2; j > 1; j-- {
				if !big.NewInt(j).ProbablyPrime(0) {
					if i+j == n {
						fmt.Println(i, j)
						return
					}
				}
			}
		}
	}
}
