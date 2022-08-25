/*
Vasya has found a strange device. On the front panel of a device there are: a red button, a blue button and a display showing some positive integer. After clicking the red button, device multiplies the displayed number by two. After clicking the blue button, device subtracts one from the number on the display. If at some point the number stops being positive, the device breaks down. The display can show arbitrarily large numbers. Initially, the display shows number _n_.

Bob wants to get number _m_ on the display. What minimum number of clicks he has to make in order to achieve this result?

Input

The first and the only line of the input contains two distinct integers _n_ and _m_ (1 ≤ _n_, _m_ ≤ 104), separated by a space .

Output

Print a single number — the minimum number of times one needs to push the button required to get the number _m_ out of number _n_.
*/

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func solve(n, m float64) float64 {
	if n == m {
		return 0
	}
	if n > m {
		return n - m
	}
	if math.Floor(m/2) != math.Ceil(m/2) {
		return solve(n, (m+1)/2) + 2
	}
	return solve(n, m/2) + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	solution := solve(float64(n), float64(m))
	writer.WriteString(strconv.FormatFloat(solution, 'f', -1, 64))
	writer.Flush()
}
