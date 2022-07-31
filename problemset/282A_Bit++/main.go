/*
The classic programming language of Bitland is Bit++. This language is so peculiar and complicated.

The language is that peculiar as it has exactly one variable, called _x_. Also, there are two operations:

- Operation ++ increases the value of variable _x_ by 1.
- Operation \-- decreases the value of variable _x_ by 1.

A statement in language Bit++ is a sequence, consisting of exactly one operation and one variable _x_. The statement is written without spaces, that is, it can only contain characters "+", "\-", "X". Executing a statement means applying the operation it contains.

A programme in Bit++ is a sequence of statements, each of them needs to be executed. Executing a programme means executing all the statements it contains.

You're given a programme in language Bit++. The initial value of _x_ is 0. Execute the programme and find its final value (the value of the variable when this programme is executed).

Input

The first line contains a single integer _n_ (1 ≤ _n_ ≤ 150) — the number of statements in the programme.

Next _n_ lines contain a statement each. Each statement contains exactly one operation (++ or \--) and exactly one variable _x_ (denoted as letter «X»). Thus, there are no empty statements. The operation and the variable can be written in any order.

Output

Print a single integer — the final value of _x_.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var x int
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		if op == "X++" || op == "++X" {
			x++
		} else {
			x--
		}
	}
	fmt.Println(x)
}
