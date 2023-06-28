package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&c)

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	fmt.Printf("Números em ordem crescente: %.2f, %.2f, %.2f\n", a, b, c)
}
