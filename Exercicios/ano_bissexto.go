package main

import (
	"fmt"
)

func AnoBissexto(ano int) string {
	if (ano%4 == 0 && ano%100 != 0) || ano%400 == 0 {
		return "Ano é bissexto!"
	} else {
		return "Ano não é bissexto"
	}
}

func main() {
	var ano int
	fmt.Print("Digite o ano: ")
	fmt.Scanln(&ano)
	fmt.Println(AnoBissexto(ano))

}
