/*
Recently, Anton has found a set. The set consists of small English letters. Anton carefully wrote out all the letters from the set in one line, separated by a comma. He also added an opening curved bracket at the beginning of the line and a closing curved bracket at the end of the line.

Unfortunately, from time to time Anton would forget writing some letter and write it again. He asks you to count the total number of distinct letters in his set.
Input

The first and the single line contains the set of letters. The length of the line doesn't exceed 1000. It is guaranteed that the line starts from an opening curved bracket and ends with a closing curved bracket. Between them, small English letters are listed, separated by a comma. Each comma is followed by a space.
Output

Print a single number — the number of distinct letters in Anton's set.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	s = strings.TrimSpace(tmp)
	var sArr []string = strings.Split(s[1:len(s)-1], ", ")
	unique := make(map[string]bool)
	for i := 0; i < len(sArr); i++ {
		unique[sArr[i]] = true
	}
	delete(unique, "")
	fmt.Println(len(unique))
}
