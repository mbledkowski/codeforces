/*
Petya loves lucky numbers. We all know that lucky numbers are the positive integers whose decimal representations contain only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

Unfortunately, not all numbers are lucky. Petya calls a number nearly lucky if the number of lucky digits in it is a lucky number. He wonders whether number _n_ is a nearly lucky number.

Input

The only line contains an integer _n_ (1 ≤ _n_ ≤ 1018).

Please do not use the %lld specificator to read or write 64-bit numbers in С++. It is preferred to use the cin, cout streams or the %I64d specificator.

Output

Print on the single line "YES" if _n_ is a nearly lucky number. Otherwise, print "NO" (without the quotes).
*/

package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	var m map[string]int = make(map[string]int)
	for i := 0; i < len(n); i++ {
		m[string(n[i])]++
	}
	var count int = m["4"] + m["7"]
	if (count == 4 || count == 7) && count != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
