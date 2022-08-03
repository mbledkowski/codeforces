/*
Limak is going to participate in a contest on the last day of the 2016. The contest will start at 20:00 and will last four hours, exactly until midnight. There will be _n_ problems, sorted by difficulty, i.e. problem 1 is the easiest and problem _n_ is the hardest. Limak knows it will take him 5·_i_ minutes to solve the _i_\-th problem.

Limak's friends organize a New Year's Eve party and Limak wants to be there at midnight or earlier. He needs _k_ minutes to get there from his house, where he will participate in the contest first.

How many problems can Limak solve if he wants to make it to the party?

Input

The only line of the input contains two integers _n_ and _k_ (1 ≤ _n_ ≤ 10, 1 ≤ _k_ ≤ 240) — the number of the problems in the contest and the number of minutes Limak needs to get to the party from his house.

Output

Print one integer, denoting the maximum possible number of problems Limak can solve so that he could get to the party at midnight or earlier.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	result := int(math.Floor(0.1 * (math.Sqrt(5)*math.Sqrt(1925-8*float64(k)) - 5)))
	if n > result {
		fmt.Println(result)
	} else {
		fmt.Println(n)
	}
}
