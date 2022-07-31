/*
An elephant decided to visit his friend. It turned out that the elephant's house is located at point 0 and his friend's house is located at point _x_(_x_ > 0) of the coordinate line. In one step the elephant can move 1, 2, 3, 4 or 5 positions forward. Determine, what is the minimum number of steps he need to make in order to get to his friend's house.

Input

The first line of the input contains an integer _x_ (1 ≤ _x_ ≤ 1 000 000) — The coordinate of the friend's house.

Output

Print the minimum number of steps that elephant needs to make to get from point 0 to point _x_.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(math.Ceil(float64(x) / 5))
}
