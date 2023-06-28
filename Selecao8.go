package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro entre 1 e 7: ")
	fmt.Scan(&numero)

	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	fmt.Printf("O número %d corresponde a %s.\n", numero, diaDaSemana)
}
