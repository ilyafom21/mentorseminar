
package main

import "fmt"

func main() {
	balance := 1000
	var choice, sum int

	for {
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Scan(&sum)
			balance += sum
		case 2:
			fmt.Scan(&sum)
			if sum > balance {
				fmt.Println("Недостаточно средств")
			} else {
				balance -= sum
			}
		case 3:
			fmt.Println(balance)
		case 4:
			return
		}
	}
}
