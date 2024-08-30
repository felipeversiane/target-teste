package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Digite uma string para reverter (ou deixe em branco para usar a string pré-definida):")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		input = "Olá, Mundo!"
	}

	reversed := reverseString(input)
	fmt.Printf("String original: %s\n", input)
	fmt.Printf("String revertida: %s\n", reversed)
}
