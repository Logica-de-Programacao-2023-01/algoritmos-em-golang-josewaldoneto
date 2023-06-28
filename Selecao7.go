package main

import "fmt"

func main() {
	var salario, novoSalario float64
	fmt.Print("Digite o salário do funcionário: ")
	fmt.Scanln(&salario)
	if salario < 1000 {
		novoSalario = salario * 1.10
	} else {
		novoSalario = salario * 1.05
	}
	fmt.Printf("O novo salário é: R$ %.2f\n", novoSalario)
}
