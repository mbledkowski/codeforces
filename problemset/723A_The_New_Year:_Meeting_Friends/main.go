/*
There are three friend living on the straight line _Ox_ in Lineland. The first friend lives at the point _x_1, the second friend lives at the point _x_2, and the third friend lives at the point _x_3. They plan to celebrate the New Year together, so they need to meet at one point. What is the minimum total distance they have to travel in order to meet at some point and celebrate the New Year?

It's guaranteed that the optimal answer is always integer.

Input

The first line of the input contains three distinct integers _x_1, _x_2 and _x_3 (1 ≤ _x_1, _x_2, _x_3 ≤ 100) — the coordinates of the houses of the first, the second and the third friends respectively.

Output

Print one integer — the minimum total distance the friends need to travel in order to meet together.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3 int
	fmt.Scan(&x1, &x2, &x3)
	fmt.Println(int(math.Abs(float64(x1-x2))+math.Abs(float64(x2-x3))+math.Abs(float64(x3-x1))) / 2)
}
