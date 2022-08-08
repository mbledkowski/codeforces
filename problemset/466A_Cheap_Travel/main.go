/*
Ann has recently started commuting by subway. We know that a one ride subway ticket costs a rubles. Besides, Ann found out that she can buy a special ticket for m rides (she can buy it several times). It costs b rubles. Ann did the math; she will need to use subway n times. Help Ann, tell her what is the minimum sum of money she will have to spend to make n rides?
Input

The single line contains four space-separated integers n, m, a, b (1 ≤ n, m, a, b ≤ 1000) — the number of rides Ann has planned, the number of rides covered by the m ride ticket, the price of a one ride ticket and the price of an m ride ticket.
Output

Print a single integer — the minimum sum in rubles that Ann will need to spend.
*/

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	pricePerTicket := float64(b) / float64(m)
	var sum float64
	if pricePerTicket < float64(a) {
		sum += math.Floor(float64(n)/float64(m)) * float64(b)
		if n%m*a > b {
			sum += float64(b)
		} else {
			sum += float64(n % m * a)
		}
	} else {
		sum += float64(n * a)
	}
	writer.WriteString(strconv.Itoa(int(sum)))
	writer.Flush()
}
