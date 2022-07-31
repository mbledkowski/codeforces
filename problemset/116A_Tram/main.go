/*
Linear Kingdom has exactly one tram line. It has _n_ stops, numbered from 1 to _n_ in the order of tram's movement. At the _i_\-th stop _a__i_ passengers exit the tram, while _b__i_ passengers enter it. The tram is empty before it arrives at the first stop. Also, when the tram arrives at the last stop, all passengers exit so that it becomes empty.

Your task is to calculate the tram's minimum capacity such that the number of people inside the tram at any time never exceeds this capacity. Note that at each stop all exiting passengers exit before any entering passenger enters the tram.

Input

The first line contains a single number _n_ (2 ≤ _n_ ≤ 1000) — the number of the tram's stops.

Then _n_ lines follow, each contains two integers _a__i_ and _b__i_ (0 ≤ _a__i_, _b__i_ ≤ 1000) — the number of passengers that exits the tram at the _i_\-th stop, and the number of passengers that enter the tram at the _i_\-th stop. The stops are given from the first to the last stop in the order of tram's movement.

- The number of people who exit at a given stop does not exceed the total number of people in the tram immediately before it arrives at the stop. More formally, ![](https://espresso.codeforces.com/eb1e20fd58cc4373d8f42eb7e743d3ea232d1a19.png). This particularly means that _a_1 = 0.
- At the last stop, all the passengers exit the tram and it becomes empty. More formally, ![](https://espresso.codeforces.com/8cbe43ecd252bf7d670f9a3956cbe50263d2f09b.png).
- No passenger will enter the train at the last stop. That is, _b__n_ = 0.

Output

Print a single integer denoting the minimum possible capacity of the tram (0 is allowed).
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a, b int
	var total, peakTotal int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		total += b - a
		if total > peakTotal {
			peakTotal = total
		}
	}

	fmt.Println(peakTotal)
}
