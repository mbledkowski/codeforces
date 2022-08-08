/*
Sereja and Dima play a game. The rules of the game are very simple. The players have n cards in a row. Each card contains a number, all numbers on the cards are distinct. The players take turns, Sereja moves first. During his turn a player can take one card: either the leftmost card in a row, or the rightmost one. The game ends when there is no more cards. The player who has the maximum sum of numbers on his cards by the end of the game, wins.

Sereja and Dima are being greedy. Each of them chooses the card with the larger number during his move.

Inna is a friend of Sereja and Dima. She knows which strategy the guys are using, so she wants to determine the final score, given the initial state of the game. Help her.
Input

The first line contains integer n (1 ≤ n ≤ 1000) — the number of cards on the table. The second line contains space-separated numbers on the cards from left to right. The numbers on the cards are distinct integers from 1 to 1000.
Output

On a single line, print two integers. The first number is the number of Sereja's points at the end of the game, the second number is the number of Dima's points at the end of the game.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func f(cards []int, move string) []int {
	result := make([]int, 2)
	if move == "Sereja" {
		if len(cards) == 1 {
			result[0] = cards[0]
			return result
		}
		if cards[0] > cards[len(cards)-1] {
			subf := f(cards[1:], "Dima")
			result[0] = cards[0] + subf[0]
			result[1] = subf[1]
		} else {
			subf := f(cards[0:len(cards)-1], "Dima")
			result[0] = cards[len(cards)-1] + subf[0]
			result[1] = subf[1]
		}
	} else {
		if len(cards) == 1 {
			result[1] = cards[0]
			return result
		}
		if cards[0] > cards[len(cards)-1] {
			subf := f(cards[1:], "Sereja")
			result[0] = subf[0]
			result[1] = cards[0] + subf[1]
		} else {
			subf := f(cards[0:len(cards)-1], "Sereja")
			result[0] = subf[0]
			result[1] = cards[len(cards)-1] + subf[1]
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	cards := make([]int, n)
	scanner.Scan()
	temp := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		card, _ := strconv.Atoi(temp[i])
		cards[i] = card
	}

	subf := f(cards, "Sereja")
	writer.WriteString(strconv.Itoa(subf[0]) + " " + strconv.Itoa(subf[1]))
	writer.Flush()
}
