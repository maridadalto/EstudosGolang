package main

import "fmt"

func fatorial(num int) int {
	var fatorial int
	for i := 1; num >= i; i++ {
		fatorial = num * i
	}
	return fatorial
}

func main() {
	var numero int
	fmt.Print("Digite um número para calcular o fatorial: ")
	fmt.Scanln(&numero)

	calcular := fatorial(numero)

	fmt.Println("O fatorial de ", numero, " é: ", calcular)
}
