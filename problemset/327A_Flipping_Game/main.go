/*
Iahub got bored, so he invented a game to be played on paper.

He writes _n_ integers _a_1, _a_2, ..., _a__n_. Each of those integers can be either 0 or 1. He's allowed to do exactly one move: he chooses two indices _i_ and _j_ (1 ≤ _i_ ≤ _j_ ≤ _n_) and flips all values _a__k_ for which their positions are in range \[_i_, _j_\] (that is _i_ ≤ _k_ ≤ _j_). Flip the value of _x_ means to apply operation _x_ = 1 - _x_.

The goal of the game is that after exactly one move to obtain the maximum number of ones. Write a program to solve the little game of Iahub.

Input

The first line of the input contains an integer _n_ (1 ≤ _n_ ≤ 100). In the second line of the input there are _n_ integers: _a_1, _a_2, ..., _a__n_. It is guaranteed that each of those _n_ values is either 0 or 1.

Output

Print an integer — the maximal number of 1s that can be obtained after exactly one move.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func countOnes(a []int) int {
	var count int
	for _, v := range a {
		if v == 1 {
			count++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	aOnes := countOnes(a)
	aZeros := n - aOnes
	max := aOnes
	if aZeros == 0 {
		max = aOnes - 1
	} else if aOnes == 0 {
		max = aZeros
	} else {
		for l := 0; l < n; l++ {
			for r := l + 1; r <= n; r++ {
				tempOnes := countOnes(a[l:r])
				tempZeros := r - l - tempOnes
				if tempZeros+aOnes-tempOnes > max {
					max = tempZeros + aOnes - tempOnes
				}
			}
		}
	}
	writer.WriteString(strconv.Itoa(max))
	writer.Flush()
}
