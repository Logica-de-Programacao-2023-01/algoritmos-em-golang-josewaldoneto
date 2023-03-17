package main

import "fmt"

// Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.

func main() {
	var preco float64
	fmt.Print("Qual o preço do produto em reais? ")
	fmt.Scan(&preco)

	desconto := preco * 0.9

	fmt.Print("O valor desse produto com um desconto de 10% será de R$", desconto)
}
