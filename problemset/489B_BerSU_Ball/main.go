/*
The Berland State University is hosting a ballroom dance in celebration of its 100500-th anniversary! n boys and m girls are already busy rehearsing waltz, minuet, polonaise and quadrille moves.

We know that several boy&girl pairs are going to be invited to the ball. However, the partners' dancing skill in each pair must differ by at most one.

For each boy, we know his dancing skills. Similarly, for each girl we know her dancing skills. Write a code that can determine the largest possible number of pairs that can be formed from n boys and m girls.
Input

The first line contains an integer n (1 ≤ n ≤ 100) — the number of boys. The second line contains sequence a1, a2, ..., an (1 ≤ ai ≤ 100), where ai is the i-th boy's dancing skill.

Similarly, the third line contains an integer m (1 ≤ m ≤ 100) — the number of girls. The fourth line contains sequence b1, b2, ..., bm (1 ≤ bj ≤ 100), where bj is the j-th girl's dancing skill.
Output

Print a single number — the required maximum possible number of pairs.
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	b := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	sort.Ints(a)
	sort.Ints(b)

	var result int
	var i, j, jmin int
	for i < n && j < m {
		if a[i]-b[j] <= 1 && b[j]-a[i] <= 1 {
			result++
			i++
			j++
			jmin = j
			continue
		}
		if j < m-1 {
			j++
		} else {
			i++
			j = jmin
		}
	}
	writer.WriteString(strconv.Itoa(result))
	writer.Flush()
}
