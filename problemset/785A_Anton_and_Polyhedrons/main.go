/*
Anton's favourite geometric figures are regular polyhedrons. Note that there are five kinds of regular polyhedrons:

- Tetrahedron. Tetrahedron has 4 triangular faces.
- Cube. Cube has 6 square faces.
- Octahedron. Octahedron has 8 triangular faces.
- Dodecahedron. Dodecahedron has 12 pentagonal faces.
- Icosahedron. Icosahedron has 20 triangular faces.

All five kinds of polyhedrons are shown on the picture below:

![](https://espresso.codeforces.com/fd444445876e0f8f0f9f9564366f27ea30053c08.png)

Anton has a collection of _n_ polyhedrons. One day he decided to know, how many faces his polyhedrons have in total. Help Anton and find this number!

Input

The first line of the input contains a single integer _n_ (1 ≤ _n_ ≤ 200 000) — the number of polyhedrons in Anton's collection.

Each of the following _n_ lines of the input contains a string _s__i_ — the name of the _i_\-th polyhedron in Anton's collection. The string can look like this:

- "Tetrahedron" (without quotes), if the _i_\-th polyhedron in Anton's collection is a tetrahedron.
- "Cube" (without quotes), if the _i_\-th polyhedron in Anton's collection is a cube.
- "Octahedron" (without quotes), if the _i_\-th polyhedron in Anton's collection is an octahedron.
- "Dodecahedron" (without quotes), if the _i_\-th polyhedron in Anton's collection is a dodecahedron.
- "Icosahedron" (without quotes), if the _i_\-th polyhedron in Anton's collection is an icosahedron.

Output

Output one number — the total number of faces in all the polyhedrons in Anton's collection.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tempn, _ := reader.ReadString('\n')
	tempn = strings.TrimSpace(tempn)
	n, _ := strconv.Atoi(tempn)

	var tetrahedron, cube, octahedron, dodecahedron, icosahedron int
	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "Tetrahedron" {
			tetrahedron++
		} else if s == "Cube" {
			cube++
		} else if s == "Octahedron" {
			octahedron++
		} else if s == "Dodecahedron" {
			dodecahedron++
		} else if s == "Icosahedron" {
			icosahedron++
		}
	}
	fmt.Println(tetrahedron*4 + cube*6 + octahedron*8 + dodecahedron*12 + icosahedron*20)
}
