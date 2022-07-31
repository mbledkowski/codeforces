/*
A soldier wants to buy _w_ bananas in the shop. He has to pay _k_ dollars for the first banana, 2_k_ dollars for the second one and so on (in other words, he has to pay _i_·_k_ dollars for the _i_\-th banana).

He has _n_ dollars. How many dollars does he have to borrow from his friend soldier to buy _w_ bananas?

Input

The first line contains three positive integers _k_, _n_, _w_ (1  ≤  _k_, _w_  ≤  1000, 0 ≤ _n_ ≤ 109), the cost of the first banana, initial number of dollars the soldier has and number of bananas he wants.

Output

Output one integer — the amount of dollars that the soldier must borrow from his friend. If he doesn't have to borrow money, output 0.
*/

package main

import "fmt"

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	var moneyPerBanana float64 = float64(k+k*w) / 2
	var moneyNeeded float64 = moneyPerBanana*float64(w) - float64(n)
	if moneyNeeded < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(int(moneyNeeded))
	}
}
