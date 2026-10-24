
package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)

	positive := 0
	negative := 0
	zero := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zero++
		}
	}

	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(zero)
}
