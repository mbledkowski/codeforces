/*
Mishka is a little polar bear. As known, little bears loves spending their free time playing dice for chocolates. Once in a wonderful sunny morning, walking around blocks of ice, Mishka met her friend Chris, and they started playing the game.

Rules of the game are very simple: at first number of rounds _n_ is defined. In every round each of the players throws a cubical dice with distinct numbers from 1 to 6 written on its faces. Player, whose value after throwing the dice is greater, wins the round. In case if player dice values are equal, no one of them is a winner.

In average, player, who won most of the rounds, is the winner of the game. In case if two players won the same number of rounds, the result of the game is draw.

Mishka is still very little and can't count wins and losses, so she asked you to watch their game and determine its result. Please help her!

Input

The first line of the input contains single integer _n_ _n_ (1 ≤ _n_ ≤ 100) — the number of game rounds.

The next _n_ lines contains rounds description. _i_\-th of them contains pair of integers _m__i_ and _c__i_ (1 ≤ _m__i_,  _c__i_ ≤ 6) — values on dice upper face after Mishka's and Chris' throws in _i_\-th round respectively.

Output

If Mishka is the winner of the game, print "Mishka" (without quotes) in the only line.

If Chris is the winner of the game, print "Chris" (without quotes) in the only line.

If the result of the game is draw, print "Friendship is magic!^^" (without quotes) in the only line.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var m, c int
	for i := 0; i < n; i++ {
		scanner.Scan()
		mici := strings.Split(scanner.Text(), " ")
		mi, _ := strconv.Atoi(mici[0])
		ci, _ := strconv.Atoi(mici[1])
		if mi > ci {
			m++
		} else if ci > mi {
			c++
		}
	}
	if m > c {
		writer.WriteString("Mishka")
	} else if c > m {
		writer.WriteString("Chris")
	} else {
		writer.WriteString("Friendship is magic!^^")
	}
	writer.Flush()
}
