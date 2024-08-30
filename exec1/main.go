package main

import "fmt"

var (
	index int
	soma  int
	k     int
)

func main() {
	index = 13

	for k < index {
		k = k + 1
		soma = soma + k
	}

	fmt.Printf("Soma: %d\n", soma)
}
