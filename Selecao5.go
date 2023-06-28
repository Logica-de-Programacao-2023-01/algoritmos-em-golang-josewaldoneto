package main

import "fmt"

func main() {
	var a int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&a)
	if a%3 == 0 && a%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e de 5.")
	} else {
		fmt.Println("O número não é múltiplo de 3 e de 5.")
	}
}
