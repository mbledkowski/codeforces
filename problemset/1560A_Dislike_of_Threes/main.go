/*
Polycarp doesn't like integers that are divisible by 3 or end with the digit 3

in their decimal representation. Integers that meet both conditions are disliked by Polycarp, too.

Polycarp starts to write out the positive (greater than 0
) integers which he likes: 1,2,4,5,7,8,10,11,14,16,…. Output the k-th element of this sequence (the elements are numbered from 1

).
Input

The first line contains one integer t
(1≤t≤100) — the number of test cases. Then t

test cases follow.

Each test case consists of one line containing one integer k
(1≤k≤1000

).
Output

For each test case, output in a separate line one integer x
— the k-th element of the sequence that was written out by Polycarp.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func f(k int) int {
	if k == 1 {
		return 1
	}
	fk := f(k - 1)
	for i := fk + 1; i < fk+4; i++ {
		if i%3 != 0 && i%10 != 3 {
			return i
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	var t int
	scanner.Scan()
	t, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		writer.WriteString(strconv.Itoa(f(k)) + "\n")
	}
	writer.Flush()
}
