/*
In Berland it is the holiday of equality. In honor of the holiday the king decided to equalize the welfare of all citizens in Berland by the expense of the state treasury.

Totally in Berland there are _n_ citizens, the welfare of each of them is estimated as the integer in _a__i_ burles (burle is the currency in Berland).

You are the royal treasurer, which needs to count the minimum charges of the kingdom on the king's present. The king can only give money, he hasn't a power to take away them.

Input

The first line contains the integer _n_ (1 ≤ _n_ ≤ 100) — the number of citizens in the kingdom.

The second line contains _n_ integers _a_1, _a_2, ..., _a__n_, where _a__i_ (0 ≤ _a__i_ ≤ 106) — the welfare of the _i_\-th citizen.

Output

In the only line print the integer _S_ — the minimum number of burles which are had to spend.
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
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	tmp = strings.TrimSpace(tmp)
	n, _ := strconv.Atoi(tmp)
	var max int
	tmp, _ = reader.ReadString('\n')
	a := strings.Split(tmp, " ")
	aint := make([]int, n)
	for i := 0; i < n; i++ {
		aint[i], _ = strconv.Atoi(strings.TrimSpace(a[i]))
		if aint[i] > max {
			max = aint[i]
		}
	}
	var sum int
	for i := 0; i < n; i++ {
		sum += max - aint[i]
	}
	fmt.Println(sum)
}
