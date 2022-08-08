/*
Vanya walks late at night along a straight street of length l, lit by n lanterns. Consider the coordinate system with the beginning of the street corresponding to the point 0, and its end corresponding to the point l. Then the i-th lantern is at the point ai. The lantern lights all points of the street that are at the distance of at most d from it, where d is some positive number, common for all lanterns.

Vanya wonders: what is the minimum light radius d should the lanterns have to light the whole street?
Input

The first line contains two integers n, l (1 ≤ n ≤ 1000, 1 ≤ l ≤ 109) — the number of lanterns and the length of the street respectively.

The next line contains n integers ai (0 ≤ ai ≤ l). Multiple lanterns can be located at the same point. The lanterns may be located at the ends of the street.
Output

Print the minimum light radius d, needed to light the whole street. The answer will be considered correct if its absolute or relative error doesn't exceed 10 - 9.
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	n, _ := strconv.Atoi(tokens[0])
	l, _ := strconv.Atoi(tokens[1])
	scanner.Scan()
	a := strings.Split(scanner.Text(), " ")
	aint := make([]int, n)
	for i := 0; i < n; i++ {
		aint[i], _ = strconv.Atoi(a[i])
	}
	sort.Ints(aint)
	d := float64(aint[0])
	for i := 1; i < n; i++ {
		locald := float64(aint[i]-aint[i-1]) / 2
		if locald > d {
			d = locald
		}
	}
	locald := float64(l - aint[n-1])
	if locald > d {
		d = locald
	}
	writer.WriteString(strconv.FormatFloat(d, 'f', -1, 64))
	writer.Flush()
}
