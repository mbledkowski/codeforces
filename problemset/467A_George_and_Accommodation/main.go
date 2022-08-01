/*
George has recently entered the BSUCP (Berland State University for Cool Programmers). George has a friend Alex who has also entered the university. Now they are moving into a dormitory.

George and Alex want to live in the same room. The dormitory has _n_ rooms in total. At the moment the _i_\-th room has _p__i_ people living in it and the room can accommodate _q__i_ people in total (_p__i_ ≤ _q__i_). Your task is to count how many rooms has free place for both George and Alex.

Input

The first line contains a single integer _n_ (1 ≤ _n_ ≤ 100) — the number of rooms.

The _i_\-th of the next _n_ lines contains two integers _p__i_ and _q__i_ (0 ≤ _p__i_ ≤ _q__i_ ≤ 100) — the number of people who already live in the _i_\-th room and the room's capacity.

Output

Print a single integer — the number of rooms where George and Alex can move in.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var p, q int
	var count int
	for i := 0; i < n; i++ {
		fmt.Scan(&p, &q)
		if p+1 < q {
			count++
		}
	}
	fmt.Println(count)
}
