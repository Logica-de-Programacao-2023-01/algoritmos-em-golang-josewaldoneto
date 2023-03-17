package main

import "fmt"

// Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

func main() {
	var idade int
	fmt.Print("Qual a sua idade em anos? ")
	fmt.Scan(&idade)

	dias := idade * 365

	fmt.Print("Sua idade é, aproximadamente, ", dias)
	fmt.Print(" dias.")
}
