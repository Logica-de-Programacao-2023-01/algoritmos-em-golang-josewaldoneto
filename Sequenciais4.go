package main

import "fmt"

// Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

func main() {
	var x1, x2, x3 float64
	fmt.Print("Escreva o número que terá peso 2. ")
	fmt.Scan(&x1)
	fmt.Print("Escreva o número que terá peso 3. ")
	fmt.Scan(&x2)
	fmt.Print("Escreva o número que terá peso 5. ")
	fmt.Scan(&x3)

	peso2 := x1 * 2
	peso3 := x2 * 3
	peso5 := x3 * 5
	media := (peso2 + peso3 + peso5) / 10

	fmt.Println("A média ponderada entre esses números é", media)
}
