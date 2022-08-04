/*
The Saratov State University Olympiad Programmers Training Center (SSU OPTC) has _n_ students. For each student you know the number of times he/she has participated in the ACM ICPC world programming championship. According to the ACM ICPC rules, each person can participate in the world championship at most 5 times.

The head of the SSU OPTC is recently gathering teams to participate in the world championship. Each team must consist of exactly three people, at that, any person cannot be a member of two or more teams. What maximum number of teams can the head make if he wants each team to participate in the world championship with the same members at least _k_ times?

Input

The first line contains two integers, _n_ and _k_ (1 ≤ _n_ ≤ 2000; 1 ≤ _k_ ≤ 5). The next line contains _n_ integers: _y_1, _y_2, ..., _y__n_ (0 ≤ _y__i_ ≤ 5), where _y__i_ shows the number of times the _i_\-th person participated in the ACM ICPC world championship.

Output

Print a single number — the answer to the problem.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	nk := strings.Split(line, " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	y := strings.Split(line, " ")
	var count int
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(y[i])
		if x+k <= 5 {
			count++
		}
	}
	fmt.Println(math.Floor(float64(count) / 3))
}
