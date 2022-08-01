/*
A word or a sentence in some language is called a pangram if all the characters of the alphabet of this language appear in it at least once. Pangrams are often used to demonstrate fonts in printing or test the output devices.

You are given a string consisting of lowercase and uppercase Latin letters. Check whether this string is a pangram. We say that the string contains a letter of the Latin alphabet if this letter occurs in the string in uppercase or lowercase.

Input

The first line contains a single integer _n_ (1 ≤ _n_ ≤ 100) — the number of characters in the string.

The second line contains the string. The string consists only of uppercase and lowercase Latin letters.

Output

Output "YES", if the string is a pangram and "NO" otherwise.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	s = strings.ToLower(s)
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	var count int
	for _, c := range alphabet {
		if strings.Contains(s, string(c)) {
			count++
		}
	}
	if count == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
