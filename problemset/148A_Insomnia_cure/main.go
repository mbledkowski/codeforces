/*
«One dragon. Two dragon. Three dragon», — the princess was counting. She had trouble falling asleep, and she got bored of counting lambs when she was nine.

However, just counting dragons was boring as well, so she entertained herself at best she could. Tonight she imagined that all dragons were here to steal her, and she was fighting them off. Every k-th dragon got punched in the face with a frying pan. Every l-th dragon got his tail shut into the balcony door. Every m-th dragon got his paws trampled with sharp heels. Finally, she threatened every n-th dragon to call her mom, and he withdrew in panic.

How many imaginary dragons suffered moral or physical damage tonight, if the princess counted a total of d dragons?
Input

Input data contains integer numbers k, l, m, n and d, each number in a separate line (1 ≤ k, l, m, n ≤ 10, 1 ≤ d ≤ 105).
Output

Output the number of damaged dragons.
*/

package main

import "fmt"

func main() {
	klmn := make([]uint8, 4)
	fmt.Scan(&klmn[0], &klmn[1], &klmn[2], &klmn[3])
	var d uint32
	fmt.Scan(&d)
	var array = make([]bool, d+1)
	for i := 0; i < 4; i++ {
		for j := uint32(klmn[i]); j <= d; j += uint32(klmn[i]) {
			array[j] = true
		}
	}
	var count uint32
	for _, e := range array {
		if e == true {
			count++
		}
	}
	fmt.Println(count)
}
