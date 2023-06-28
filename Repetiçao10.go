package main

import "fmt"

func main() {
	var numero, maior int
	fmt.Println("Digite uma sequência de números inteiros (para parar, digite 0):")
	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		if numero > maior {
			maior = numero
		}
	}
	if maior == 0 {
		fmt.Println("Não foi digitado nenhum número diferente de zero")
	} else {
		fmt.Printf("O maior número digitado foi %d\n", maior)
	}
}
