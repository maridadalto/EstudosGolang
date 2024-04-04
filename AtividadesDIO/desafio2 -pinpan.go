package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PIN e PAN")
			continue
		} else if i%3 == 0 {
			fmt.Println("PIN")
			continue
		} else if i%5 == 0 {
			fmt.Println("PAN")
			continue
		} else {
			fmt.Println(i)
		}

	}
}
