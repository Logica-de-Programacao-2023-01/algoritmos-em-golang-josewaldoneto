package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.

func main() {
	var numero int
	fmt.Print("Escreva um número.")
	fmt.Scan(&numero)

	dobro := numero * 2
	triplo := numero * 3
	quadruplo := numero * 4

	fmt.Print("O dobro deste número é ", dobro)

	fmt.Print(", o triplo deste número é ", triplo)

	fmt.Print(" e o quadruplo deste número é ", quadruplo)
}
