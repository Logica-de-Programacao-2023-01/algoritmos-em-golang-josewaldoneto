package main

import "fmt"

// Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.

func main() {
	var salary float64
	fmt.Print("Qual é o seu salário em reais? ")
	fmt.Scan(&salary)

	aumento := salary * 1.15

	fmt.Print("Seu novo salário com um aumento de 15% é R$", aumento)
}
