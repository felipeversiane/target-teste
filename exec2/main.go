package main

import (
	"fmt"
)

func main() {
	number := 21

	if isFibonacci(number) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", number)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", number)
	}
}

func isFibonacci(n int) bool {
	if n < 0 {
		return false
	}

	a, b := 0, 1
	for a <= n {
		if a == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}
