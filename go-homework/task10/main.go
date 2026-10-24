
package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	fmt.Scan(&x)

	sum := x
	min := x
	max := x

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		sum += x

		if x < min {
			min = x
		}

		if x > max {
			max = x
		}
	}

	average := float64(sum) / float64(n)

	fmt.Println(sum)
	fmt.Println(average)
	fmt.Println(min)
	fmt.Println(max)
}
