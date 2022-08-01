/*
The translation from the Berland language into the Birland language is not an easy task. Those languages are very similar: a berlandish word differs from a birlandish word with the same meaning a little: it is spelled (and pronounced) reversely. For example, a Berlandish word code corresponds to a Birlandish word edoc. However, it's easy to make a mistake during the «translation». Vasya translated word _s_ from Berlandish into Birlandish as _t_. Help him: find out if he translated the word correctly.

Input

The first line contains word _s_, the second line contains word _t_. The words consist of lowercase Latin letters. The input data do not consist unnecessary spaces. The words are not empty and their lengths do not exceed 100 symbols.

Output

If the word _t_ is a word _s_, written reversely, print YES, otherwise print NO.
*/

package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if len(s) != len(t) {
		fmt.Println("NO")
		return
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[len(t)-i-1] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
