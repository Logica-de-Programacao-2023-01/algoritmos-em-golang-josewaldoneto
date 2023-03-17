package main

import "fmt"

// Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).

func main() {
	var altura, peso float64
	fmt.Print("Qual é a sua altura em metros?")
	fmt.Scan(&altura)
	fmt.Print("Qual o seu peso em quilogramas?")
	fmt.Scan(&peso)

	imc := peso / (altura * altura)

	fmt.Println("Seu IMC é", imc)
}
