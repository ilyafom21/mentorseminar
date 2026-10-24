
package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	fmt.Scan(&x)
	min, max := x, x

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println(max)
	fmt.Println(min)
}
