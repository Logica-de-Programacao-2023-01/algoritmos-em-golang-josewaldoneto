package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scanln(&c)
	if a < b && a < c {
		fmt.Printf("O menor número é: %d\n", a)
	} else if b < c {
		fmt.Printf("O menor número é: %d\n", b)
	} else {
		fmt.Printf("O menor número é: %d\n", c)
	}
}
