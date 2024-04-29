package main

import (
	"fmt"
)

func verificarLetra(palavra string, letra string) int {
	var qtde int

	for _, caractere := range palavra {
		if string(caractere) == letra {
			qtde++
		}
	}
	return qtde
}

func main() {
	var palavra, letra string

	fmt.Print("Insira a palavra a ser comparada: ")
	fmt.Scanln(&palavra)
	fmt.Print("Insira a letra a ser comparada: ")
	fmt.Scanln(&letra)

	qtde_letras := verificarLetra(palavra, letra)
	fmt.Printf("A quantidade da letra [%s] na palavra [%s] é: %d\n", letra, palavra, qtde_letras)

}
