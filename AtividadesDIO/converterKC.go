package main

import "fmt"

const K float64 = 373.0

var C float64

func main() {

	C = K - 273.0

	fmt.Printf("O valor do ponto de ebulição em Kelvin é: %.2f e o ponto de ebulição em Celsius é: %.2f", K, C)

}
