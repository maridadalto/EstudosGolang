package main

import "fmt"

func ReverseString(palavra string) string {
	runes := []rune(palavra)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func main() {
	var palavra string
	fmt.Print("Informe a palavra a ser invertida: ")
	fmt.Scanln(&palavra)
	fmt.Println("Palavra invertida:", ReverseString(palavra))
}
