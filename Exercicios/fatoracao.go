package main

import "fmt"

func fatoracao(numero int) {
	fmt.Printf("Fatoração de %d é: ", numero)
	divisor := 2

	for numero > 1 {
		if numero%divisor == 0 {
			fmt.Printf("%d ", divisor)
			numero /= divisor
		} else {
			divisor++
		}
	}
	fmt.Println()
}

func main() {
	var numero int
	fmt.Print("Digite um número para calcular a fatoração: ")
	fmt.Scanln(&numero)

	fatoracao(numero)
}
