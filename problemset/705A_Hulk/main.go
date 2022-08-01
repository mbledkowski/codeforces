/*
Dr. Bruce Banner hates his enemies (like others don't). As we all know, he can barely talk when he turns into the incredible Hulk. That's why he asked you to help him to express his feelings.

Hulk likes the Inception so much, and like that his feelings are complicated. They have _n_ layers. The first layer is hate, second one is love, third one is hate and so on...

For example if _n_ = 1, then his feeling is "I hate it" or if _n_ = 2 it's "I hate that I love it", and if _n_ = 3 it's "I hate that I love that I hate it" and so on.

Please help Dr. Banner.

Input

The only line of the input contains a single integer _n_ (1 ≤ _n_ ≤ 100) — the number of layers of love and hate.

Output

Print Dr.Banner's feeling in one line.
*/

package main

import "fmt"

func hulk(n int) {
	if n == 0 {
		return
	}
	hulk(n - 1)
	if n%2 == 0 {
		fmt.Print("I love that ")
	} else {
		fmt.Print("I hate that ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	hulk(n - 1)
	if n%2 == 0 {
		fmt.Print("I love it")
	} else {
		fmt.Print("I hate it")
	}
}
