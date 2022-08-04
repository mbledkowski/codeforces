/*
The School №0 of the capital of Berland has _n_ children studying in it. All the children in this school are gifted: some of them are good at programming, some are good at maths, others are good at PE (Physical Education). Hence, for each child we know value _t__i_:

- _t__i_ = 1, if the _i_\-th child is good at programming,
- _t__i_ = 2, if the _i_\-th child is good at maths,
- _t__i_ = 3, if the _i_\-th child is good at PE

Each child happens to be good at exactly one of these three subjects.

The Team Scientific Decathlon Olympias requires teams of three students. The school teachers decided that the teams will be composed of three children that are good at different subjects. That is, each team must have one mathematician, one programmer and one sportsman. Of course, each child can be a member of no more than one team.

What is the maximum number of teams that the school will be able to present at the Olympiad? How should the teams be formed for that?

Input

The first line contains integer _n_ (1 ≤ _n_ ≤ 5000) — the number of children in the school. The second line contains _n_ integers _t_1, _t_2, ..., _t__n_ (1 ≤ _t__i_ ≤ 3), where _t__i_ describes the skill of the _i_\-th child.

Output

In the first line output integer _w_ — the largest possible number of teams.

Then print _w_ lines, containing three numbers in each line. Each triple represents the indexes of the children forming the team. You can print both the teams, and the numbers in the triplets in any order. The children are numbered from 1 to _n_ in the order of their appearance in the input. Each child must participate in no more than one team. If there are several solutions, print any of them.

If no teams can be compiled, print the only line with value _w_ equal to 0.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solve(n int, t []int, writer *bufio.Writer) {
	var a, b, c []int
	for i := 0; i < n; i++ {
		switch t[i] {
		case 1:
			a = append(a, i+1)
		case 2:
			b = append(b, i+1)
		case 3:
			c = append(c, i+1)
		}
	}
	w := len(a)
	if len(b) < w {
		w = len(b)
	}
	if len(c) < w {
		w = len(c)
	}

	writer.WriteString(strconv.Itoa(w))
	for i := 0; i < w; i++ {
		writer.WriteString("\n")
		writer.WriteString(strconv.Itoa(a[i]))
		writer.WriteString(" ")
		writer.WriteString(strconv.Itoa(b[i]))
		writer.WriteString(" ")
		writer.WriteString(strconv.Itoa(c[i]))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	writer := bufio.NewWriter(os.Stdout)
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	t := strings.Split(scanner.Text(), " ")
	tint := make([]int, n)
	for i := 0; i < n; i++ {
		tint[i], _ = strconv.Atoi(t[i])
	}
	solve(n, tint, writer)
	writer.Flush()
}
