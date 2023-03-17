package main

import "fmt"

// Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.

func main() {
	var x1, x2, x3 int
	fmt.Print("Escreva o primeiro número inteiro que deseja somar.")
	fmt.Scan(&x1)
	fmt.Print("Escreva o segundo número inteiro que deseja somar.")
	fmt.Scan(&x2)
	fmt.Print("Escreva o terceiro número inteiro que deseja somar.")
	fmt.Scan(&x3)

	soma := x1 + x2 + x3

	fmt.Print("A soma desses três números é ", soma)
}
