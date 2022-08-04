/*
You are given two integers a and b

.

In one move, you can choose some integer k
from 1 to 10 and add it to a or subtract it from a. In other words, you choose an integer k∈[1;10] and perform a:=a+k or a:=a−k. You may use different values of k

in different moves.

Your task is to find the minimum number of moves required to obtain b
from a

.

You have to answer t

independent test cases.
Input

The first line of the input contains one integer t
(1≤t≤2⋅104) — the number of test cases. Then t

test cases follow.

The only line of the test case contains two integers a
and b (1≤a,b≤109

).
Output

For each test case, print the answer: the minimum number of moves required to obtain b
from a.
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		diff := int(math.Abs(float64(a - b)))
		var moves int
		if diff%10 == 0 {
			moves = diff / 10
		} else {
			moves = int(math.Floor(float64(diff)/10)) + 1
		}
		fmt.Println(moves)
	}
}
