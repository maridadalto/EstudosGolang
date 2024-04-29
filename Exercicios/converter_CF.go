package main

import "fmt"

func ConverterCelsius(temp float64) float64 {
	fahrenheit := (temp * 9 / 5) + 32
	return fahrenheit
}

func main() {
	var celsius float64
	fmt.Print("Informe a temperatura em Celsius para conversão em Fahrenheit: ")
	fmt.Scanln(&celsius)
	resultado := ConverterCelsius(celsius)
	fmt.Printf("O resultado da conversão:  %v°C para %v°F\n", celsius, resultado)

}
