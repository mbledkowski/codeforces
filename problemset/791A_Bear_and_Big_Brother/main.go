/*
Bear Limak wants to become the largest of bears, or at least to become larger than his brother Bob.

Right now, Limak and Bob weigh _a_ and _b_ respectively. It's guaranteed that Limak's weight is smaller than or equal to his brother's weight.

Limak eats a lot and his weight is tripled after every year, while Bob's weight is doubled after every year.

After how many full years will Limak become strictly larger (strictly heavier) than Bob?

Input

The only line of the input contains two integers _a_ and _b_ (1 ≤ _a_ ≤ _b_ ≤ 10) — the weight of Limak and the weight of Bob respectively.

Output

Print one integer, denoting the integer number of years after which Limak will become strictly larger than Bob.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var firstLog, secondLog float64 = math.Log(float64(a) / float64(b)), math.Log(float64(3)) - math.Log(float64(2))
	var result float64 = -firstLog / secondLog
	var resultInt float64 = math.Ceil(result)
	if result == resultInt {
		resultInt++
	}
	fmt.Println(resultInt)
}
