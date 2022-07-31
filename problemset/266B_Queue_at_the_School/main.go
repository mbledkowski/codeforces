/*
During the break the schoolchildren, boys and girls, formed a queue of _n_ people in the canteen. Initially the children stood in the order they entered the canteen. However, after a while the boys started feeling awkward for standing in front of the girls in the queue and they started letting the girls move forward each second.

Let's describe the process more precisely. Let's say that the positions in the queue are sequentially numbered by integers from 1 to _n_, at that the person in the position number 1 is served first. Then, if at time _x_ a boy stands on the _i_\-th position and a girl stands on the (_i_ + 1)\-th position, then at time _x_ + 1 the _i_\-th position will have a girl and the (_i_ + 1)\-th position will have a boy. The time is given in seconds.

You've got the initial position of the children, at the initial moment of time. Determine the way the queue is going to look after _t_ seconds.

Input

The first line contains two integers _n_ and _t_ (1 ≤ _n_, _t_ ≤ 50), which represent the number of children in the queue and the time after which the queue will transform into the arrangement you need to find.

The next line contains string _s_, which represents the schoolchildren's initial arrangement. If the _i_\-th position in the queue contains a boy, then the _i_\-th character of string _s_ equals "B", otherwise the _i_\-th character equals "G".

Output

Print string _a_, which describes the arrangement after _t_ seconds. If the _i_\-th position has a boy after the needed time, then the _i_\-th character _a_ must equal "B", otherwise it must equal "G".
*/

package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	var s string
	fmt.Scan(&s)
	a := make([]byte, n)
	a = []byte(s)

	for i := 0; i < t; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] == 'B' && a[j+1] == 'G' {
				a[j] = 'G'
				a[j+1] = 'B'
				j++
			}
		}
	}

	fmt.Println(string(a))
}
