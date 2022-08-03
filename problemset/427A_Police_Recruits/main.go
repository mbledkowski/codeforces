/*
The police department of your city has just started its journey. Initially, they don’t have any manpower. So, they started hiring new recruits in groups.

Meanwhile, crimes keeps occurring within the city. One member of the police force can investigate only one crime during his/her lifetime.

If there is no police officer free (isn't busy with crime) during the occurrence of a crime, it will go untreated.

Given the chronological order of crime occurrences and recruit hirings, find the number of crimes which will go untreated.

Input

The first line of input will contain an integer _n_ (1 ≤ _n_ ≤ 105), the number of events. The next line will contain _n_ space-separated integers.

If the integer is -1 then it means a crime has occurred. Otherwise, the integer will be positive, the number of officers recruited together at that time. No more than 10 officers will be recruited at a time.

Output

Print a single integer, the number of crimes which will go untreated.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	var arr []int = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	var count, countMax int
	for i := 0; i < n; i++ {
		num := arr[i]
		count -= num
		if count > countMax {
			countMax = count
		}
	}
	fmt.Fprintf(writer, "%d\n", countMax)

	writer.Flush()
}
