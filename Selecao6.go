package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	if num1 > 0 && num2 > 0 {
		resultado := num1 * num2
		fmt.Printf("O resultado da multiplicação é: %d\n", resultado)
	} else {
		resultado := num1 + num2
		fmt.Printf("O resultado da soma é: %d\n", resultado)
	}
}
