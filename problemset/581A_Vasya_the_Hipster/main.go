/*
One day Vasya the Hipster decided to count how many socks he had. It turned out that he had _a_ red socks and _b_ blue socks.

According to the latest fashion, hipsters should wear the socks of different colors: a red one on the left foot, a blue one on the right foot.

Every day Vasya puts on new socks in the morning and throws them away before going to bed as he doesn't want to wash them.

Vasya wonders, what is the maximum number of days when he can dress fashionable and wear different socks, and after that, for how many days he can then wear the same socks until he either runs out of socks or cannot make a single pair from the socks he's got.

Can you help him?

Input

The single line of the input contains two positive integers _a_ and _b_ (1 ≤ _a_, _b_ ≤ 100) — the number of red and blue socks that Vasya's got.

Output

Print two space-separated integers — the maximum number of days when Vasya can wear different socks and the number of days when he can wear the same socks until he either runs out of socks or cannot make a single pair from the socks he's got.

Keep in mind that at the end of the day Vasya throws away the socks that he's been wearing on that day.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println(b, math.Floor(float64(a-b)/2))
	} else {
		fmt.Println(a, math.Floor(float64(b-a)/2))
	}
}
