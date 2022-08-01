/*
There is a game called "I Wanna Be the Guy", consisting of _n_ levels. Little X and his friend Little Y are addicted to the game. Each of them wants to pass the whole game.

Little X can pass only _p_ levels of the game. And Little Y can pass only _q_ levels of the game. You are given the indices of levels Little X can pass and the indices of levels Little Y can pass. Will Little X and Little Y pass the whole game, if they cooperate each other?

Input

The first line contains a single integer _n_ (1 ≤  _n_ ≤ 100).

The next line contains an integer _p_ (0 ≤ _p_ ≤ _n_) at first, then follows _p_ distinct integers _a_1, _a_2, ..., _a__p_ (1 ≤ _a__i_ ≤ _n_). These integers denote the indices of levels Little X can pass. The next line contains the levels Little Y can pass in the same format. It's assumed that levels are numbered from 1 to _n_.

Output

If they can pass all the levels, print "I become the guy.". If it's impossible, print "Oh, my keyboard!" (without the quotes).
*/

package main

import "fmt"

func main() {
	var n, p, q int
	fmt.Scan(&n)
	a, b := make([]int, n), make([]int, n)
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&b[i])
	}
	levels := make([]bool, n)
	for i := 0; i < p; i++ {
		levels[a[i]-1] = true
	}
	for i := 0; i < q; i++ {
		levels[b[i]-1] = true
	}
	for i := 0; i < n; i++ {
		if !levels[i] {
			fmt.Println("Oh, my keyboard!")
			return
		}
	}
	fmt.Println("I become the guy.")
}
