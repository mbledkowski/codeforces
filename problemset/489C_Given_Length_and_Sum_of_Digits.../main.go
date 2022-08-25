/*
You have a positive integer _m_ and a non-negative integer _s_. Your task is to find the smallest and the largest of the numbers that have length _m_ and sum of digits _s_. The required numbers should be non-negative integers written in the decimal base without leading zeroes.

Input

The single line of the input contains a pair of integers _m_, _s_ (1 ≤ _m_ ≤ 100, 0 ≤ _s_ ≤ 900) — the length and the sum of the digits of the required numbers.

Output

In the output print the pair of the required non-negative integer numbers — first the minimum possible number, then — the maximum possible number. If no numbers satisfying conditions required exist, print the pair of numbers "\-1 -1" (without the quotes).
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	min, max := make([]int, m), make([]int, m)
	temps := s
	for i := m - 1; i > 0; i-- {
		if temps > 9 {
			temps -= 9
			min[i] = 9
		} else {
			min[i] = temps - 1
			if temps > 0 {
				temps = 1
			} else {
				temps = 0
			}
			break
		}
	}
	min[0] = temps
	temps = s
	for i := 0; i < m; i++ {
		if temps > 9 {
			temps -= 9
			max[i] = 9
		} else {
			max[i] = temps
			break
		}
	}
	if s == 0 && m == 1 {
		writer.WriteString("0 0")
	} else if min[0] == 0 || min[0] > 9 {
		writer.WriteString("-1 -1")
	} else {
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(min[i]))
		}
		writer.WriteString(" ")
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(max[i]))
		}
	}
	writer.Flush()
}
