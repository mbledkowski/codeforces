/*
Find the minimum area of a square land on which you can place two identical rectangular a×b

houses. The sides of the houses should be parallel to the sides of the desired square land.

Formally,

    You are given two identical rectangles with side lengths a

and b (1≤a,b≤100

    ) — positive integers (you are given just the sizes, but not their positions).
    Find the square of the minimum area that contains both given rectangles. Rectangles can be rotated (both or just one), moved, but the sides of the rectangles should be parallel to the sides of the desired square.

Two rectangles can touch each other (side or corner), but cannot intersect. Rectangles can also touch the sides of the square but must be completely inside it. You can rotate the rectangles. Take a look at the examples for a better understanding.
The picture shows a square that contains red and green rectangles.
Input

The first line contains an integer t
(1≤t≤10000) —the number of test cases in the input. Then t

test cases follow.

Each test case is a line containing two integers a
, b (1≤a,b≤100

) — side lengths of the rectangles.
Output

Print t
answers to the test cases. Each answer must be a single integer — minimal area of square land, that contains two rectangles with dimensions a×b.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		ab := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])
		if a > b {
			if b+b > a {
				writer.WriteString(strconv.Itoa((b*2)*(b*2)) + "\n")
			} else {
				writer.WriteString(strconv.Itoa(a*a) + "\n")
			}
		} else {
			if a+a > b {
				writer.WriteString(strconv.Itoa((a*2)*(a*2)) + "\n")
			} else {
				writer.WriteString(strconv.Itoa(b*b) + "\n")
			}
		}
	}
	writer.Flush()
}
