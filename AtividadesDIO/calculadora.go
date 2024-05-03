package main

import "fmt"

func main() {

	for {
		fmt.Println("Selecione uma opção:")
		fmt.Println("1. Soma")
		fmt.Println("2. Subtração")
		fmt.Println("3. MUltiplicação")
		fmt.Println("4. Divisão")
		fmt.Println("5. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			fmt.Printf("Resultado da soma: %d\n", Soma(10, 5))
		case 2:
			fmt.Printf("Resultado da subtração: %d\n", Subtracao(10, 5))
		case 3:
			fmt.Printf("Resultado da multiplicação: %d\n", Multiplica(10, 5))
		case 4:
			fmt.Printf("Resultado da divisão: %d\n", Divisao(10, 5))
		case 5:
			return

		}

	}
}
func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func Subtracao(i ...int) int {
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}
func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}
func Divisao(i ...int) int {
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			return 0 // Retorna 0 se encontrar divisão por zero
		}
		total /= v

	}
	return total
}
