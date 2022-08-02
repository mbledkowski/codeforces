/*
Manao works on a sports TV. He's spent much time watching the football games of some country. After a while he began to notice different patterns. For example, each team has two sets of uniforms: home uniform and guest uniform. When a team plays a game at home, the players put on the home uniform. When a team plays as a guest on somebody else's stadium, the players put on the guest uniform. The only exception to that rule is: when the home uniform color of the host team matches the guests' uniform, the host team puts on its guest uniform as well. For each team the color of the home and guest uniform is different.

There are _n_ teams taking part in the national championship. The championship consists of _n_·(_n_ - 1) games: each team invites each other team to its stadium. At this point Manao wondered: how many times during the championship is a host team going to put on the guest uniform? Note that the order of the games does not affect this number.

You know the colors of the home and guest uniform for each team. For simplicity, the colors are numbered by integers in such a way that no two distinct colors have the same number. Help Manao find the answer to his question.

Input

The first line contains an integer _n_ (2 ≤ _n_ ≤ 30). Each of the following _n_ lines contains a pair of distinct space-separated integers _h__i_, _a__i_ (1 ≤ _h__i_, _a__i_ ≤ 100) — the colors of the _i_\-th team's home and guest uniforms, respectively.

Output

In a single line print the number of games where the host team is going to play in the guest uniform.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	teams := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		team := scanner.Text()
		teams[i] = team
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if strings.Split(teams[i], " ")[0] == strings.Split(teams[j], " ")[1] {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
