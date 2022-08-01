/*
Mad scientist Mike entertains himself by arranging rows of dominoes. He doesn't need dominoes, though: he uses rectangular magnets instead. Each magnet has two poles, positive (a "plus") and negative (a "minus"). If two magnets are put together at a close distance, then the like poles will repel each other and the opposite poles will attract each other.

Mike starts by laying one magnet horizontally on the table. During each following step Mike adds one more magnet horizontally to the right end of the row. Depending on how Mike puts the magnet on the table, it is either attracted to the previous one (forming a group of multiple magnets linked together) or repelled by it (then Mike lays this magnet at some distance to the right from the previous one). We assume that a sole magnet not linked to others forms a group of its own.

![](https://espresso.codeforces.com/a2cfc00fa5ba91709a0c55ed4bb5e70f976b5f20.png)

Mike arranged multiple magnets in a row. Determine the number of groups that the magnets formed.

Input

The first line of the input contains an integer _n_ (1 ≤ _n_ ≤ 100000) — the number of magnets. Then _n_ lines follow. The _i_\-th line (1 ≤ _i_ ≤ _n_) contains either characters "01", if Mike put the _i_\-th magnet in the "plus-minus" position, or characters "10", if Mike put the magnet in the "minus-plus" position.

Output

On the single line of the output print the number of groups of magnets.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var s, t byte
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(tmp))

	count := n

	for i := 0; i < n; i++ {
		tmp, _ = reader.ReadString('\n')
		s = tmp[0]

		if s == t {
			count--
		} else {
			t = s
		}
	}

	fmt.Println(count)
}
