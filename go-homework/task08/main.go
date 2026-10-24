
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	prime := true

	if n < 2 {
		prime = false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println("Простое")
	} else {
		fmt.Println("Не простое")
	}
}
