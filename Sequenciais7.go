package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.

func main() {
	var numero int
	fmt.Print("Escreva um número. ")
	fmt.Scan(&numero)

	x := numero
	y := numero

	x++
	fmt.Println("O sucessor desse número é", x)
	y--
	fmt.Println("O antecessor desse número é", y)
}
