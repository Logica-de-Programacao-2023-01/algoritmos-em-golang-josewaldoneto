package main

import "fmt"

// Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.

func main() {
	var dias, diaria float64
	fmt.Print("Escreva o número de dias trabalhados. ")
	fmt.Scan(&dias)
	fmt.Print("Escreva o valor da sua diária em reais. ")
	fmt.Scan(&diaria)

	salario := dias * diaria

	fmt.Println("Seu salário será de R$", salario)
}
