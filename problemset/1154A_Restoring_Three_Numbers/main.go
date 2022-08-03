/*
Polycarp has guessed three positive integers a, b and c. He keeps these numbers in secret, but he writes down four numbers on a board in arbitrary order — their pairwise sums (three numbers) and sum of all three numbers (one number). So, there are four numbers on a board in random order: a+b, a+c, b+c and a+b+c

.

You have to guess three numbers a
, b and c

using given numbers. Print three guessed integers in any order.

Pay attention that some given numbers a
, b and c can be equal (it is also possible that a=b=c

).
Input

The only line of the input contains four positive integers x1,x2,x3,x4
(2≤xi≤109) — numbers written on a board in random order. It is guaranteed that the answer exists for the given number x1,x2,x3,x4

.
Output

Print such positive integers a
, b and c that four numbers written on a board are values a+b, a+c, b+c and a+b+c written in some order. Print a, b and c in any order. If there are several answers, you can print any. It is guaranteed that the answer exists.
*/

package main

import "fmt"

func main() {
	var x []int = make([]int, 4)
	fmt.Scan(&x[0], &x[1], &x[2], &x[3])
	var max, maxi int
	for i, v := range x {
		if v > max {
			max = v
			maxi = i
		}
	}

	var i1, i2, i3 int
	if maxi == 0 {
		i1, i2, i3 = 1, 2, 3
	} else if maxi == 1 {
		i1, i2, i3 = 0, 2, 3
	} else if maxi == 2 {
		i1, i2, i3 = 0, 1, 3
	} else {
		i1, i2, i3 = 0, 1, 2
	}

	a := (x[i1] + x[i2] - x[i3]) / 2
	b := x[i1] - a
	c := x[i2] - a

	fmt.Println(a, b, c)
}
