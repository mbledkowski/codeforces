/*
Little Vasya loves orange juice very much. That's why any food and drink in his kitchen necessarily contains orange juice. There are _n_ drinks in his fridge, the volume fraction of orange juice in the _i_\-th drink equals _p__i_ percent.

One day Vasya decided to make himself an orange cocktail. He took equal proportions of each of the _n_ drinks and mixed them. Then he wondered, how much orange juice the cocktail has.

Find the volume fraction of orange juice in the final drink.

Input

The first input line contains a single integer _n_ (1 ≤ _n_ ≤ 100) — the number of orange-containing drinks in Vasya's fridge. The second line contains _n_ integers _p__i_ (0 ≤ _p__i_ ≤ 100) — the volume fraction of orange juice in the _i_\-th drink, in percent. The numbers are separated by a space.

Output

Print the volume fraction in percent of orange juice in Vasya's cocktail. The answer will be considered correct if the absolute or relative error does not exceed 10  - 4.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var sum int
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		sum += p
	}
	fmt.Printf("%f\n", float64(sum)/float64(n))
}
