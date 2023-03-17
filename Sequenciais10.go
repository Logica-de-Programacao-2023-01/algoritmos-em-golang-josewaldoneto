package main

import "fmt"

// Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.

func main() {
	var peso float64
	fmt.Print("Escreva seu peso em quilogramas. ")
	fmt.Scan(&peso)

	pesolibras := peso * 2.205

	fmt.Print("Seu peso em libras será ", pesolibras)
	fmt.Print(",aproximadamente.")
}
