/*
You are living on an infinite plane with the Cartesian coordinate system on it. In one move you can go to any of the four adjacent points (left, right, up, down).

More formally, if you are standing at the point (x,y)

, you can:

    go left, and move to (x−1,y)

, or
go right, and move to (x+1,y)
, or
go up, and move to (x,y+1)
, or
go down, and move to (x,y−1)

    .

There are n
boxes on this plane. The i-th box has coordinates (xi,yi). It is guaranteed that the boxes are either on the x-axis or the y-axis. That is, either xi=0 or yi=0

.

You can collect a box if you and the box are at the same point. Find the minimum number of moves you have to perform to collect all of these boxes if you have to start and finish at the point (0,0)

.
Input

The first line contains a single integer t
(1≤t≤100

) — the number of test cases.

The first line of each test case contains a single integer n
(1≤n≤100

) — the number of boxes.

The i
-th line of the following n lines contains two integers xi and yi (−100≤xi,yi≤100) — the coordinate of the i-th box. It is guaranteed that either xi=0 or yi=0

.

Do note that the sum of n

over all test cases is not bounded.
Output

For each test case output a single integer — the minimum number of moves required.
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
		n, _ := strconv.Atoi(scanner.Text())
		var numOfMoves int
		var minx, miny, maxx, maxy int
		for j := 0; j < n; j++ {
			scanner.Scan()
			xystr := strings.Split(scanner.Text(), " ")
			x, _ := strconv.Atoi(xystr[0])
			y, _ := strconv.Atoi(xystr[1])
			if x < minx {
				minx = x
			} else if x > maxx {
				maxx = x
			} else if y < miny {
				miny = y
			} else if y > maxy {
				maxy = y
			}
		}
		numOfMoves += -minx*2 - miny*2 + maxx*2 + maxy*2
		writer.WriteString(strconv.Itoa(numOfMoves) + "\n")
	}
	writer.Flush()
}
