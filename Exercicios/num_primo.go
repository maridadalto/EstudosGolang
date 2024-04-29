package main

import (
	"fmt"
	"math"
)

func NumPrimo(numero int) bool {
	if numero <= 1 {
		return false
	} else if numero <= 3 {
		return true
	} else if numero%2 == 0 {
		return false
	} else {
		for i := 3; i <= int(math.Sqrt(float64(numero))); i += 2 {
			if numero%i == 0 {
				return false
			}
		}
	}
	return true
}
func main() {
	var numero int
	fmt.Print("Informe o numero para verificar se e primo:")
	fmt.Scanln(&numero)
	if NumPrimo(numero) {
		fmt.Println("E numero primo\n")
	} else {
		fmt.Println("Nao e primo\n")
	}
}
