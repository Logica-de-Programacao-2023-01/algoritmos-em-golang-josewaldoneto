package main

// Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)
	if a > b {
		fmt.Printf("O maior número é: %d\n", a)
	} else {
		fmt.Printf("O maior número é: %d\n", b)
	}
}
