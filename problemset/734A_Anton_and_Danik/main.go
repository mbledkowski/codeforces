/*
Anton likes to play chess, and so does his friend Danik.

Once they have played _n_ games in a row. For each game it's known who was the winner — Anton or Danik. None of the games ended with a tie.

Now Anton wonders, who won more games, he or Danik? Help him determine this.

Input

The first line of the input contains a single integer _n_ (1 ≤ _n_ ≤ 100 000) — the number of games played.

The second line contains a string _s_, consisting of _n_ uppercase English letters 'A' and 'D' — the outcome of each of the games. The _i_\-th character of the string is equal to 'A' if the Anton won the _i_\-th game and 'D' if Danik won the _i_\-th game.

Output

If Anton won more games than Danik, print "Anton" (without quotes) in the only line of the output.

If Danik won more games than Anton, print "Danik" (without quotes) in the only line of the output.

If Anton and Danik won the same number of games, print "Friendship" (without quotes).
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	var anton, danik int
	for _, v := range s {
		if v == 'A' {
			anton++
		} else {
			danik++
		}
	}

	if anton > danik {
		fmt.Println("Anton")
	} else if danik > anton {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
