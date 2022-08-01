/*
A Ministry for Defense sent a general to inspect the Super Secret Military Squad under the command of the Colonel SuperDuper. Having learned the news, the colonel ordered to all _n_ squad soldiers to line up on the parade ground.

By the military charter the soldiers should stand in the order of non-increasing of their height. But as there's virtually no time to do that, the soldiers lined up in the arbitrary order. However, the general is rather short-sighted and he thinks that the soldiers lined up correctly if the first soldier in the line has the maximum height and the last soldier has the minimum height. Please note that the way other solders are positioned does not matter, including the case when there are several soldiers whose height is maximum or minimum. Only the heights of the first and the last soldier are important.

For example, the general considers the sequence of heights (4, 3, 4, 2, 1, 1) correct and the sequence (4, 3, 1, 2, 2) wrong.

Within one second the colonel can swap any two neighboring soldiers. Help him count the minimum time needed to form a line-up which the general will consider correct.

Input

The first input line contains the only integer _n_ (2 ≤ _n_ ≤ 100) which represents the number of soldiers in the line. The second line contains integers _a_1, _a_2, ..., _a__n_ (1 ≤ _a__i_ ≤ 100) the values of the soldiers' heights in the order of soldiers' heights' increasing in the order from the beginning of the line to its end. The numbers are space-separated. Numbers _a_1, _a_2, ..., _a__n_ are not necessarily different.

Output

Print the only integer — the minimum number of seconds the colonel will need to form a line-up the general will like.
*/

package main

import "fmt"

func solve(n uint8, a []uint8) uint8 {
	min, max := uint8(101), uint8(1)
	var mini, maxi uint8
	for i := uint8(0); i < n; i++ {
		if a[i] <= min {
			min = a[i]
			mini = i
		}
		if a[i] > max {
			max = a[i]
			maxi = i
		}
	}
	if mini > maxi {
		return maxi + n - uint8(1) - mini
	} else {
		return maxi + n - uint8(2) - mini
	}
}

func main() {
	var n uint8
	fmt.Scan(&n)
	a := make([]uint8, n)
	for i := uint8(0); i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(solve(n, a))
}
